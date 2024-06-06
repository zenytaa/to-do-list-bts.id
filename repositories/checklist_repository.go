package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"to-do-list-bts.id/entities"
)

type ChecklistRepoOpts struct {
	Db *sql.DB
}

type ChecklistRepository interface {
	CreateOneChecklist(ctx context.Context, checkList entities.Cheklist) error
	FindAllChecklist(ctx context.Context) ([]entities.Cheklist, error)
}

type ChecklistRepositoryImpl struct {
	db *sql.DB
}

func NewChecklistRepository(chROpts *ChecklistRepoOpts) ChecklistRepository {
	return &ChecklistRepositoryImpl{
		db: chROpts.Db,
	}
}

func (r *ChecklistRepositoryImpl) CreateOneChecklist(ctx context.Context, checkList entities.Cheklist) error {
	fmt.Println(checkList)
	err := r.db.QueryRowContext(ctx, qCreateOneChecklist, checkList.Name).Scan(&checkList.Id)
	if err != nil {
		return err
	}
	return nil
}

func (r *ChecklistRepositoryImpl) FindAllChecklist(ctx context.Context) ([]entities.Cheklist, error) {
	checklists := []entities.Cheklist{}

	rows, err := r.db.QueryContext(ctx, qFindAllChecklist)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		checklist := entities.Cheklist{}
		err := rows.Scan(&checklist.Id, &checklist.Name)
		if err != nil {
			return nil, err
		}
		checklists = append(checklists, checklist)
	}

	return checklists, nil
}
