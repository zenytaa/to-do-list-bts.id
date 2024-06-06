package usecases

import (
	"context"

	"to-do-list-bts.id/entities"
	"to-do-list-bts.id/repositories"
)

type ItemUsecaseOpts struct {
	ItemRepo      repositories.ItemRepository
	ChecklistRepo repositories.ChecklistRepository
}

type ItemUsecase interface {
	CreateItem(ctx context.Context, item entities.Item) error
}

type ItemUsecaseImpl struct {
	ItemRepository      repositories.ItemRepository
	ChecklistRepository repositories.ChecklistRepository
}

func NewItemUsecaseImpl(chiUOpts *ItemUsecaseOpts) ItemUsecase {
	return &ItemUsecaseImpl{
		ItemRepository:      chiUOpts.ItemRepo,
		ChecklistRepository: chiUOpts.ChecklistRepo,
	}
}

func (u *ItemUsecaseImpl) CreateItem(ctx context.Context, item entities.Item) error {
	_, err := u.ChecklistRepository.FindOneById(ctx, item.Checklist.Id)
	if err != nil {
		return err
	}

	err = u.ItemRepository.CreateOne(ctx, item)
	if err != nil {
		return err
	}
	return nil
}
