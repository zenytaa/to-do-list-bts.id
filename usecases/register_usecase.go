package usecases

import (
	"context"

	"to-do-list-bts.id/entities"
	"to-do-list-bts.id/repositories"
	"to-do-list-bts.id/utils"
)

type RegisterUsecaseOpts struct {
	HashAlgorithm     utils.Hasher
	AuthTokenProvider utils.AuthTokenProvider
	UserRepo          repositories.UserRepository
}

type RegisterUsecase interface {
	RegisterUser(ctx context.Context, user entities.User) error
}

type RegisterUsecaseImpl struct {
	HashAlgorithm     utils.Hasher
	AuthTokenProvider utils.AuthTokenProvider
	UserRepository    repositories.UserRepository
}

func NewRegisterUsecaseImpl(registerOpts *RegisterUsecaseOpts) RegisterUsecase {
	return &RegisterUsecaseImpl{
		HashAlgorithm:     registerOpts.HashAlgorithm,
		AuthTokenProvider: registerOpts.AuthTokenProvider,
		UserRepository:    registerOpts.UserRepo,
	}
}

func (u *RegisterUsecaseImpl) RegisterUser(ctx context.Context, user entities.User) error {
	pwd := user.Password
	pwdHash, err := u.HashAlgorithm.HashPassword(pwd.String)
	if err != nil {
		return err
	}

	user.Password = *utils.ByteToNullString(pwdHash)

	_, err = u.UserRepository.CreateOneUser(ctx, user)
	if err != nil {
		return err
	}

	return nil
}
