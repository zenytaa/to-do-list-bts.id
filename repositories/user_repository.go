package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
	"to-do-list-bts.id/constants"
	"to-do-list-bts.id/custom_errors"
	"to-do-list-bts.id/entities"
)

type UserRepoOpt struct {
	Db *sql.DB
}

type UserRepository interface {
	CreateOneUser(ctx context.Context, u entities.User) (*entities.User, error)
	FindOneByUsername(ctx context.Context, username string) (*entities.User, error)
}

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepositoryPostgres(trOpt *UserRepoOpt) UserRepository {
	return &UserRepositoryImpl{
		db: trOpt.Db,
	}
}

func (r *UserRepositoryImpl) CreateOneUser(ctx context.Context, u entities.User) (*entities.User, error) {
	newU := entities.User{}

	values := []interface{}{}
	values = append(values, u.Name)

	if u.Password.Valid {
		values = append(values, u.Password)
	}

	err := r.db.QueryRowContext(ctx, qCreateOneUser, values...).Scan(&newU.Id)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == constants.ViolatesUniqueConstraintPgErrCode {
			return nil, custom_errors.BadRequest(err, constants.UserEmailNotUniqueErrMsg)
		}
		return nil, err
	}

	return &newU, nil
}

func (r *UserRepositoryImpl) FindOneByUsername(ctx context.Context, username string) (*entities.User, error) {
	u := entities.User{}

	err := r.db.QueryRowContext(ctx, qFindUserByUsername, username).Scan(&u.Id, &u.Name, &u.Password)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, custom_errors.NotFound(err)
		}
		return nil, err
	}

	return &u, nil
}
