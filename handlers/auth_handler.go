package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"to-do-list-bts.id/constants"
	"to-do-list-bts.id/dtos"
	"to-do-list-bts.id/entities"
	"to-do-list-bts.id/usecases"
	"to-do-list-bts.id/utils"
)

type AuthHandlerOpts struct {
	LoginUsecase    usecases.LoginUsecase
	RegisterUsecase usecases.RegisterUsecase
}

type AuthHandler struct {
	LoginUsecase    usecases.LoginUsecase
	RegisterUsecase usecases.RegisterUsecase
}

func NewAuthHandler(ahOpts *AuthHandlerOpts) *AuthHandler {
	return &AuthHandler{
		LoginUsecase:    ahOpts.LoginUsecase,
		RegisterUsecase: ahOpts.RegisterUsecase,
	}
}

func (h *AuthHandler) Login(ctx *gin.Context) {
	var payload dtos.LoginRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		_ = ctx.Error(err)
		return
	}

	token, err := h.LoginUsecase.LoginUser(ctx, payload.Username, payload.Password)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dtos.ResponseMessage{
		Message: constants.ResponseMsgLogin,
		Data:    token,
	})
}

func (h *AuthHandler) RegisterUser(ctx *gin.Context) {
	var payload dtos.RegisterRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		_ = ctx.Error(err)
		return
	}

	u := entities.User{
		Name:     payload.Username,
		Email:    payload.Email,
		Password: *utils.StringToNullString(*payload.Password),
	}

	err := h.RegisterUsecase.RegisterUser(ctx, u)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dtos.ResponseMessage{
		Message: constants.ResponseMsgRegistered,
		Data:    nil,
	})
}
