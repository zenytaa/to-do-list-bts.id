package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"to-do-list-bts.id/constants"
	"to-do-list-bts.id/dtos"
)

func RequestId(c *gin.Context) {
	uuid, err := uuid.NewUUID()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, dtos.ErrResponse{
			Message: constants.ResponseMsgErrorInternalServer,
		})
		return
	}

	c.Set(constants.RequestId, uuid)
	c.Next()
}
