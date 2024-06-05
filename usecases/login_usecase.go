package usecases

import (
	"context"

	"to-do-list-bts.id/constants"
	"to-do-list-bts.id/custom_errors"
	"to-do-list-bts.id/repositories"
	"to-do-list-bts.id/utils"
)

type LoginUsecaseOpts struct {
	UserRepo          repositories.UserRepository
	HashAlgorithm     utils.Hasher
	AuthTokenProvider utils.AuthTokenProvider
}

type LoginUsecase interface {
	LoginUser(ctx context.Context, email, password string) (*string, error)
}

type LoginUsecaseImpl struct {
	UserRepository    repositories.UserRepository
	HashAlgorithm     utils.Hasher
	AuthTokenProvider utils.AuthTokenProvider
}

func NewLoginUsecaseImpl(loginOpts *LoginUsecaseOpts) LoginUsecase {
	return &LoginUsecaseImpl{
		UserRepository:    loginOpts.UserRepo,
		HashAlgorithm:     loginOpts.HashAlgorithm,
		AuthTokenProvider: loginOpts.AuthTokenProvider,
	}
}

func (u *LoginUsecaseImpl) LoginUser(ctx context.Context, username string, password string) (*string, error) {
	user, err := u.UserRepository.FindOneByUsername(ctx, username)
	if err != nil {
		return nil, custom_errors.Unauthorized(err, constants.InvalidCredentialsErrMsg)
	}

	isCorrectPassword, err := u.HashAlgorithm.CheckPassword(password, []byte(user.Password.String))
	if !isCorrectPassword {
		return nil, custom_errors.Unauthorized(err, constants.InvalidCredentialsErrMsg)
	}

	dataTokenMap := make(map[string]interface{})
	dataTokenMap["id"] = user.Id

	accessToken, err := u.AuthTokenProvider.CreateAndSign(dataTokenMap)
	if err != nil {
		return nil, err
	}

	return &accessToken, nil
}
