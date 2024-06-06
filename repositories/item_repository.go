package repositories

import (
	"context"
	"database/sql"

	"to-do-list-bts.id/entities"
)

type ItemRepoOpts struct {
	Db *sql.DB
}

type ItemRepository interface {
	CreateOne(ctx context.Context, Item entities.Item) error
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
	err := r.db.QueryRowContext(ctx, qCreateItem).Scan(item.Id)
	if err != nil {
		return err
	}
	return nil
}
