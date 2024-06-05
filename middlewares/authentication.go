package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"to-do-list-bts.id/constants"
	"to-do-list-bts.id/custom_errors"
	"to-do-list-bts.id/dtos"
	"to-do-list-bts.id/utils"
)

func JwtAuthMiddleware(config utils.Config) func(*gin.Context) {
	return func(ctx *gin.Context) {
		authorized, data, err := utils.NewJwtProvider(config).IsAuthorized(ctx)
		if !authorized && err != nil && data == nil {
			if err.Error() == custom_errors.TokenExpired().Error() {
				ctx.AbortWithStatusJSON(http.StatusForbidden, dtos.ErrResponse{
					Message: constants.ExpiredTokenErrMsg,
				})
				return
			}
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, dtos.ErrResponse{
				Message: constants.ResponseMsgUnauthorized,
			})
			return
		}
		ctx.Set("data", data)
		ctx.Next()
	}
}
