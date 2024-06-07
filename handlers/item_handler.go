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

func (h *ItemHandler) GetAllItem(ctx *gin.Context) {
	checklistIdStr := ctx.Param("checklistId")
	checklistId, err := strconv.Atoi(checklistIdStr)
	if err != nil {
		ctx.Error(err)
		return
	}

	items, err := h.ItemUsecase.GetAllItem(ctx, int64(checklistId))
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dtos.ResponseMessage{
		Message: constants.ResponseMsgOK,
		Data:    dtos.ConvertToItemResponses(items),
	})
}

func (h *ItemHandler) GetItemById(ctx *gin.Context) {
	checklistIdStr := ctx.Param("checklistId")
	checklistId, err := strconv.Atoi(checklistIdStr)
	if err != nil {
		ctx.Error(err)
		return
	}

	itemIdStr := ctx.Param("checklistItemId")
	itemId, err := strconv.Atoi(itemIdStr)
	if err != nil {
		ctx.Error(err)
		return
	}

	item, err := h.ItemUsecase.GetOneItem(ctx, entities.Item{Id: int64(itemId), Checklist: entities.Checklist{Id: int64(checklistId)}})
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dtos.ResponseMessage{
		Message: constants.ResponseMsgOK,
		Data:    dtos.ConvertToItemResponse(item),
	})
}

func (h *ItemHandler) UpdateStatusItem(ctx *gin.Context) {
	checklistIdStr := ctx.Param("checklistId")
	checklistId, err := strconv.Atoi(checklistIdStr)
	if err != nil {
		ctx.Error(err)
		return
	}

	itemIdStr := ctx.Param("checklistItemId")
	itemId, err := strconv.Atoi(itemIdStr)
	if err != nil {
		ctx.Error(err)
		return
	}

	err = h.ItemUsecase.UpdateItemStatus(ctx, entities.Item{Id: int64(itemId), Checklist: entities.Checklist{Id: int64(checklistId)}})
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dtos.ResponseMessage{
		Message: constants.ResponseMsgUpdated,
		Data:    nil,
	})
}

func (h *ItemHandler) DeleteItem(ctx *gin.Context) {
	checklistIdStr := ctx.Param("checklistId")
	checklistId, err := strconv.Atoi(checklistIdStr)
	if err != nil {
		ctx.Error(err)
		return
	}

	itemIdStr := ctx.Param("checklistItemId")
	itemId, err := strconv.Atoi(itemIdStr)
	if err != nil {
		ctx.Error(err)
		return
	}

	err = h.ItemUsecase.DeleteItem(ctx, entities.Item{Id: int64(itemId), Checklist: entities.Checklist{Id: int64(checklistId)}})
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dtos.ResponseMessage{
		Message: constants.ResponseMsgDeleted,
		Data:    nil,
	})
}

func (h *ItemHandler) UpdateItemName(ctx *gin.Context) {
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

	itemIdStr := ctx.Param("checklistItemId")
	itemId, err := strconv.Atoi(itemIdStr)
	if err != nil {
		ctx.Error(err)
		return
	}

	item := entities.Item{
		Id:        int64(itemId),
		ItemName:  payload.ItemName,
		Checklist: entities.Checklist{Id: int64(checklistId)},
	}

	err = h.ItemUsecase.UpdateItemName(ctx, item)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dtos.ResponseMessage{
		Message: constants.ResponseMsgUpdated,
		Data:    nil,
	})
}
