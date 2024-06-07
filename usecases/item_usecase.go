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
	GetAllItem(ctx context.Context, checklistId int64) ([]entities.Item, error)
	GetOneItem(ctx context.Context, item entities.Item) (*entities.Item, error)
	UpdateItemStatus(ctx context.Context, item entities.Item) error
	DeleteItem(ctx context.Context, item entities.Item) error
	UpdateItemName(ctx context.Context, item entities.Item) error
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

func (u *ItemUsecaseImpl) GetAllItem(ctx context.Context, checklistId int64) ([]entities.Item, error) {
	_, err := u.ChecklistRepository.FindOneById(ctx, checklistId)
	if err != nil {
		return nil, err
	}

	items, err := u.ItemRepository.FindAllByChecklistId(ctx, checklistId)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (u *ItemUsecaseImpl) GetOneItem(ctx context.Context, item entities.Item) (*entities.Item, error) {
	_, err := u.ChecklistRepository.FindOneById(ctx, item.Checklist.Id)
	if err != nil {
		return nil, err
	}

	getItem, err := u.ItemRepository.FindOneByItemId(ctx, item)
	if err != nil {
		return nil, err
	}

	return getItem, nil
}

func (u *ItemUsecaseImpl) UpdateItemStatus(ctx context.Context, item entities.Item) error {
	_, err := u.ChecklistRepository.FindOneById(ctx, item.Checklist.Id)
	if err != nil {
		return err
	}

	err = u.ItemRepository.UpdateItemStatusById(ctx, item)
	if err != nil {
		return err
	}

	return nil
}

func (u *ItemUsecaseImpl) DeleteItem(ctx context.Context, item entities.Item) error {
	_, err := u.ChecklistRepository.FindOneById(ctx, item.Checklist.Id)
	if err != nil {
		return err
	}

	err = u.ItemRepository.DeleteOneItem(ctx, item)
	if err != nil {
		return err
	}

	return nil
}

func (u *ItemUsecaseImpl) UpdateItemName(ctx context.Context, item entities.Item) error {
	_, err := u.ChecklistRepository.FindOneById(ctx, item.Checklist.Id)
	if err != nil {
		return err
	}

	err = u.ItemRepository.UpdateItemName(ctx, item)
	if err != nil {
		return err
	}

	return nil
}
