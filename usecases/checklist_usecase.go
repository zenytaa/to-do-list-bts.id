package usecases

import (
	"context"

	"to-do-list-bts.id/entities"
	"to-do-list-bts.id/repositories"
)

type ChecklistUsecaseOpts struct {
	ChecklistRepo repositories.ChecklistRepository
}

type ChecklistUsecas interface {
	CreateChecklist(ctx context.Context, checklist entities.Cheklist) error
	GetAllChecklist(ctx context.Context) ([]entities.Cheklist, error)
}

type ChecklistUsecaseImpl struct {
	ChecklistRepository repositories.ChecklistRepository
}

func NewChecklistUsecaseImpl(chUOpts *ChecklistUsecaseOpts) ChecklistUsecas {
	return &ChecklistUsecaseImpl{
		ChecklistRepository: chUOpts.ChecklistRepo,
	}
}

func (u *ChecklistUsecaseImpl) CreateChecklist(ctx context.Context, checklist entities.Cheklist) error {
	err := u.ChecklistRepository.CreateOneChecklist(ctx, checklist)
	if err != nil {
		return err
	}
	return nil
}

func (u *ChecklistUsecaseImpl) GetAllChecklist(ctx context.Context) ([]entities.Cheklist, error) {
	checklists, err := u.ChecklistRepository.FindAllChecklist(ctx)
	if err != nil {
		return nil, err
	}
	return checklists, nil
}