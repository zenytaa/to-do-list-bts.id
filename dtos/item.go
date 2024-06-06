package dtos

type ItemRequest struct {
	ItemName string `json:"itemName" binding:"required"`
}
