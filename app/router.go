package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"to-do-list-bts.id/constants"
	"to-do-list-bts.id/dtos"
	"to-do-list-bts.id/handlers"
	"to-do-list-bts.id/middlewares"
	"to-do-list-bts.id/repositories"
	"to-do-list-bts.id/usecases"
	"to-do-list-bts.id/utils"
)

type RouterOpt struct {
	AuthHandler      *handlers.AuthHandler
	ChecklistHandler *handlers.ChecklistHandler
}

func createRouter(config utils.Config) *gin.Engine {
	db, err := ConnectDB(config)
	if err != nil {
		log.Fatalf("error connecting to DB: %s", err.Error())
	}

	userRepo := repositories.NewUserRepositoryPostgres(&repositories.UserRepoOpt{Db: db})
	checklistRepo := repositories.NewChecklistRepository(&repositories.ChecklistRepoOpts{Db: db})

	loginUsecase := usecases.NewLoginUsecaseImpl(&usecases.LoginUsecaseOpts{
		UserRepo:          userRepo,
		HashAlgorithm:     utils.NewBCryptHasher(),
		AuthTokenProvider: utils.NewJwtProvider(config),
	})
	registerUsecase := usecases.NewRegisterUsecaseImpl(&usecases.RegisterUsecaseOpts{
		HashAlgorithm:     utils.NewBCryptHasher(),
		AuthTokenProvider: utils.NewJwtProvider(config),
		UserRepo:          userRepo,
	})
	checklistUsecase := usecases.NewChecklistUsecaseImpl(&usecases.ChecklistUsecaseOpts{
		ChecklistRepo: checklistRepo,
	})

	authHandler := handlers.NewAuthHandler(&handlers.AuthHandlerOpts{
		LoginUsecase:    loginUsecase,
		RegisterUsecase: registerUsecase,
	})
	checklistHandler := handlers.NewChecklistHandler(&handlers.CheklistHandlerOpts{
		ChecklistUsecase: checklistUsecase,
	})

	return NewRouter(config, &RouterOpt{
		AuthHandler:      authHandler,
		ChecklistHandler: checklistHandler,
	})
}

func NewRouter(config utils.Config, handlers *RouterOpt) *gin.Engine {
	router := gin.Default()

	router.ContextWithFallback = true

	router.Use(middlewares.CORS, middlewares.RequestId, middlewares.ErrorHandling)

	publicRouter := router.Group("/api")
	publicRouter.POST("/login", handlers.AuthHandler.Login)
	publicRouter.POST("/register", handlers.AuthHandler.RegisterUser)

	privateRouter := router.Group("/api")
	{
		privateRouter.Use(middlewares.JwtAuthMiddleware(config))
		privateRouter.POST("/checklist", handlers.ChecklistHandler.CreateChecklist)
		privateRouter.GET("/checklist", handlers.ChecklistHandler.GetAllChecklist)
		privateRouter.DELETE("/checklist/:id", handlers.ChecklistHandler.DeleteChecklist)
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, dtos.ErrResponse{Message: constants.EndpointNotFoundErrMsg})
	})

	return router
}

func Init() {
	config, err := utils.ConfigInit()
	if err != nil {
		log.Fatalf("error getting env: %s", err.Error())
	}

	router := createRouter(config)

	srv := http.Server{
		Handler: router,
		Addr:    fmt.Sprintf(":%s", config.Port),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 3)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	go func() {
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Fatal("Server Shutdown: ", err)
		}
	}()

	<-ctx.Done()
	log.Println("Server exiting")

}
