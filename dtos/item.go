package dtos

import "to-do-list-bts.id/entities"

type ItemRequest struct {
	ItemName string `json:"itemName" binding:"required"`
}

type ItemResponse struct {
	Id       int64  `json:"id"`
	ItemName string `json:"item_name"`
	IsDone   bool   `json:"is_done"`
}

type ItemResponses struct {
	ChecklistId int64          `json:"checklist_id"`
	Items       []ItemResponse `json:"items"`
}

func ConvertToItemResponse(item *entities.Item) *ItemResponse {
	return &ItemResponse{
		Id:       item.Id,
		ItemName: item.ItemName,
		IsDone:   item.IsDone,
	}
}

func ConvertToItemResponses(items []entities.Item) *ItemResponses {
	responses := []ItemResponse{}

	for _, item := range items {
		responses = append(responses, *ConvertToItemResponse(&item))
	}

	return &ItemResponses{
		ChecklistId: items[0].Checklist.Id,
		Items:       responses,
	}
}
