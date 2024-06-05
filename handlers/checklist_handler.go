package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"to-do-list-bts.id/constants"
	"to-do-list-bts.id/dtos"
	"to-do-list-bts.id/entities"
	"to-do-list-bts.id/usecases"
)

type CheklistHandlerOpts struct {
	ChecklistUsecas usecases.ChecklistUsecas
}

type ChecklistHandler struct {
	ChecklistUsecas usecases.ChecklistUsecas
}

func NewChecklistHandler(chHOpts *CheklistHandlerOpts) *ChecklistHandler {
	return &ChecklistHandler{
		ChecklistUsecas: chHOpts.ChecklistUsecas,
	}
}

func (h *ChecklistHandler) CreateChecklist(ctx *gin.Context) {
	var payload dtos.ChecklistRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		_ = ctx.Error(err)
		return
	}

	checklist := entities.Cheklist{
		Name: payload.Name,
	}

	err := h.ChecklistUsecas.CreateChecklist(ctx, checklist)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dtos.ResponseMessage{
		Message: constants.ResponseMsgCreatedChecklist,
		Data:    nil,
	})
}

func (h *ChecklistHandler) GetAllChecklist(ctx *gin.Context) {
	checklists, err := h.ChecklistUsecas.GetAllChecklist(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dtos.ResponseMessage{
		Message: constants.ResponseMsgOK,
		Data:    dtos.ConvertToChecklisResponses(checklists),
	})
}
