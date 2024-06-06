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

type CheklistHandlerOpts struct {
	ChecklistUsecase usecases.ChecklistUsecase
}

type ChecklistHandler struct {
	ChecklistUsecase usecases.ChecklistUsecase
}

func NewChecklistHandler(chHOpts *CheklistHandlerOpts) *ChecklistHandler {
	return &ChecklistHandler{
		ChecklistUsecase: chHOpts.ChecklistUsecase,
	}
}

func (h *ChecklistHandler) CreateChecklist(ctx *gin.Context) {
	var payload dtos.ChecklistRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		_ = ctx.Error(err)
		return
	}

	checklist := entities.Checklist{
		Name: payload.Name,
	}

	err := h.ChecklistUsecase.CreateChecklist(ctx, checklist)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dtos.ResponseMessage{
		Message: constants.ResponseMsgCreated,
		Data:    nil,
	})
}

func (h *ChecklistHandler) GetAllChecklist(ctx *gin.Context) {
	checklists, err := h.ChecklistUsecase.GetAllChecklist(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dtos.ResponseMessage{
		Message: constants.ResponseMsgOK,
		Data:    dtos.ConvertToChecklisResponses(checklists),
	})
}

func (h *ChecklistHandler) DeleteChecklist(ctx *gin.Context) {
	checklistIdStr := ctx.Param("checklistId")
	checklistId, err := strconv.Atoi(checklistIdStr)
	if err != nil {
		ctx.Error(err)
		return
	}

	err = h.ChecklistUsecase.DeleteChecklist(ctx, int64(checklistId))
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dtos.ResponseMessage{
		Message: constants.ResponseMsgDeleted,
		Data:    nil,
	})
}
