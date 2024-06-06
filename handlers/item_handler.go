package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"to-do-list-bts.id/constants"
	"to-do-list-bts.id/dtos"
	"to-do-list-bts.id/entities"
	"to-do-list-bts.id/usecases"
)

type ItemHandlerOpts struct {
	ItemUsecase usecases.ItemUsecase
}

type ItemHandler struct {
	ItemUsecase usecases.ItemUsecase
}

func NewItemHandler(chiHOpts *ItemHandlerOpts) *ItemHandler {
	return &ItemHandler{ItemUsecase: chiHOpts.ItemUsecase}
}

func (h *ItemHandler) CreateItem(ctx *gin.Context) {
	var payload dtos.ItemRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		_ = ctx.Error(err)
		return
	}

	checklistIdStr := ctx.Param("checklistId")
	checklistId, err := strconv.Atoi(checklistIdStr)
	if err != nil {
		ctx.Error(err)
		return
	}

	item := entities.Item{
		ItemName:  payload.ItemName,
		Checklist: entities.Checklist{Id: int64(checklistId)},
		IsDone:    false,
	}

	err = h.ItemUsecase.CreateItem(ctx, item)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dtos.ResponseMessage{
		Message: constants.ResponseMsgCreated,
		Data:    nil,
	})
}
