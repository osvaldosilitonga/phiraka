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

func (u *userImpl) Login(ctx context.Context, payload web.LoginReq) (string, error) {
	user, err := u.UserRepo.FindByUsername(ctx, payload.Username)
	if err != nil {
		return "", err
	}

	isMatch := helpers.ComparePassword(user.Password, payload.Password)
	if !isMatch {
		return "", errors.New("wrong password")
	}

	// Generate Token
	token, err := helpers.GenerateToken(user.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *userImpl) Delete(ctx context.Context, username string) error {
	if err := u.UserRepo.DeleteUser(ctx, username); err != nil {
		return err
	}

	return nil
}

func (u *userImpl) FindAllUser(ctx context.Context) ([]web.FindAllUserResp, error) {
	result := []web.FindAllUserResp{}

	users, err := u.UserRepo.FindAllUser(ctx)
	if err != nil {
		return result, err
	}

	for _, v := range users {
		usr := web.FindAllUserResp{
			Username:   v.Username,
			Password:   v.Password,
			CreateTime: v.CreateTime,
		}

		result = append(result, usr)
	}

	return result, nil
}
