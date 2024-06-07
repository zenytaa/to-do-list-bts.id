package repositories

import (
	"context"
	"database/sql"

	"to-do-list-bts.id/custom_errors"
	"to-do-list-bts.id/entities"
)

type ItemRepoOpts struct {
	Db *sql.DB
}

type ItemRepository interface {
	CreateOne(ctx context.Context, Item entities.Item) error
	FindAllByChecklistId(ctx context.Context, checklistId int64) ([]entities.Item, error)
	FindOneByItemId(ctx context.Context, item entities.Item) (*entities.Item, error)
	UpdateItemStatusById(ctx context.Context, item entities.Item) error
	DeleteOneItem(ctx context.Context, item entities.Item) error
	UpdateItemName(ctx context.Context, item entities.Item) error
}

type ItemRepositoryImpl struct {
	db *sql.DB
}

func NewItemRepositoryImpl(chiROpts *ItemRepoOpts) ItemRepository {
	return &ItemRepositoryImpl{
		db: chiROpts.Db,
	}
}

func (r *ItemRepositoryImpl) CreateOne(ctx context.Context, item entities.Item) error {
	err := r.db.QueryRowContext(ctx, qCreateItem, item.ItemName, item.Checklist.Id).Scan(&item.Id)
	if err != nil {
		return err
	}
	return nil
}

func (r *ItemRepositoryImpl) FindAllByChecklistId(ctx context.Context, checklistId int64) ([]entities.Item, error) {
	items := []entities.Item{}

	rows, err := r.db.QueryContext(ctx, qFindAllItem, checklistId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		item := entities.Item{}
		err := rows.Scan(&item.Id, &item.ItemName, &item.Checklist.Id, &item.IsDone)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

func (r *ItemRepositoryImpl) FindOneByItemId(ctx context.Context, item entities.Item) (*entities.Item, error) {
	newItem := entities.Item{}

	err := r.db.QueryRowContext(ctx, qFindOneItemById, item.Id, item.Checklist.Id).Scan(&newItem.Id, &newItem.ItemName, &newItem.Checklist.Id, &newItem.IsDone)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, custom_errors.NotFound(err)
		}
		return nil, err
	}

	return &newItem, nil
}

func (r *ItemRepositoryImpl) UpdateItemStatusById(ctx context.Context, item entities.Item) error {
	stmt, err := r.db.PrepareContext(ctx, qUpdateItemStatus)
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, item.Id, item.Checklist.Id)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return custom_errors.NotFound(sql.ErrNoRows)
	}

	return nil
}

func (r *ItemRepositoryImpl) DeleteOneItem(ctx context.Context, item entities.Item) error {
	stmt, err := r.db.PrepareContext(ctx, qDeleteOneItem)
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, item.Id, item.Checklist.Id)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return custom_errors.NotFound(sql.ErrNoRows)
	}

	return nil
}

func (r *ItemRepositoryImpl) UpdateItemName(ctx context.Context, item entities.Item) error {
	stmt, err := r.db.PrepareContext(ctx, qUpdateItemName)
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, item.Id, item.Checklist.Id, item.ItemName)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return custom_errors.NotFound(sql.ErrNoRows)
	}

	return nil

}
