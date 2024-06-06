package dtos

import "to-do-list-bts.id/entities"

type ChecklistRequest struct {
	Name string `json:"name" binding:"required"`
}

type ChecklistResponse struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func ConvertToChecklisResponse(checklist entities.Checklist) *ChecklistResponse {
	return &ChecklistResponse{
		Id:   checklist.Id,
		Name: checklist.Name,
	}
}

func ConvertToChecklisResponses(checklists []entities.Checklist) []ChecklistResponse {
	response := []ChecklistResponse{}

	for _, ch := range checklists {
		response = append(response, *ConvertToChecklisResponse(ch))
	}

	return response
}
