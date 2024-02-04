package handlers

import (
	"context"
	"errors"

	"github.com/osvaldosilitonga/phiraka/server/domain/entity"
	"github.com/osvaldosilitonga/phiraka/server/domain/web"
	"github.com/osvaldosilitonga/phiraka/server/helpers"
	"github.com/osvaldosilitonga/phiraka/server/repositories"
)

type userImpl struct {
	UserRepo repositories.User
}

func NewUserHandler(ur repositories.User) User {
	return &userImpl{
		UserRepo: ur,
	}
}

func (u *userImpl) Register(ctx context.Context, payload web.RegisterReq) error {
	user := entity.User{
		Username: payload.Username,
	}

	pass, err := helpers.HashPassword(payload.Password)
	if err != nil {
		return errors.New("error while hashing password")
	}

	user.Password = pass

	err = u.UserRepo.Save(ctx, &user)
	if err != nil {
		return err
	}

	return nil
}
