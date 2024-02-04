package controllers

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/osvaldosilitonga/phiraka/server/domain/web"
	"github.com/osvaldosilitonga/phiraka/server/handlers"
	"github.com/osvaldosilitonga/phiraka/server/helpers"
	"github.com/osvaldosilitonga/phiraka/server/utils"
)

type userImpl struct {
	UserHandler handlers.User
}

func NewUserController(uh handlers.User) User {
	return &userImpl{
		UserHandler: uh,
	}
}

func (u *userImpl) Register(c echo.Context) error {
	payload := web.RegisterReq{}

	if err := c.Bind(&payload); err != nil {
		return utils.ErrorMessage(c, &utils.ApiBadRequest, err.Error())
	}
	if err := c.Validate(&payload); err != nil {
		return utils.ErrorMessage(c, &utils.ApiBadRequest, helpers.SplitError(err))
	}

	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Second*5)
	defer cancel()

	if err := u.UserHandler.Register(ctx, payload); err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return utils.ErrorMessage(c, &utils.ApiBadRequest, "username already exist")
		}

		return utils.ErrorMessage(c, &utils.ApiInternalServer, err.Error())
	}

	return utils.SuccessMessage(c, &utils.ApiCreate, nil)
}

func (u *userImpl) Login(c echo.Context) error {
	payload := web.LoginReq{}

	if err := c.Bind(&payload); err != nil {
		return utils.ErrorMessage(c, &utils.ApiBadRequest, err.Error())
	}
	if err := c.Validate(&payload); err != nil {
		return utils.ErrorMessage(c, &utils.ApiBadRequest, helpers.SplitError(err))
	}

	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Second*5)
	defer cancel()

	token, err := u.UserHandler.Login(ctx, payload)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return utils.ErrorMessage(c, &utils.ApiNotFound, err.Error())
		}
		if strings.Contains(err.Error(), "wrong password") {
			return utils.ErrorMessage(c, &utils.ApiBadRequest, err.Error())
		}

		return utils.ErrorMessage(c, &utils.ApiInternalServer, err.Error())
	}

	// Set session cookie
	cookie := new(http.Cookie)
	cookie.Name = "Authorization"
	cookie.HttpOnly = true
	cookie.Path = "/"
	cookie.Value = token
	cookie.Expires = time.Now().Add(5 * time.Hour)
	c.SetCookie(cookie)

	return utils.SuccessMessage(c, &utils.ApiOk, nil)
}

func (u *userImpl) Logout(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "Authorization"
	cookie.HttpOnly = true
	cookie.Path = "/"
	cookie.Value = ""
	cookie.SameSite = http.SameSiteLaxMode
	cookie.MaxAge = -1
	c.SetCookie(cookie)

	return utils.SuccessMessage(c, &utils.ApiOk, nil)
}

func (u *userImpl) Delete(c echo.Context) error {
	username := strings.ToLower(c.Param("username"))
	if len(username) < 4 || len(username) > 128 {
		return utils.ErrorMessage(c, &utils.ApiBadRequest, "username must be between 4 - 128 character")
	}

	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Second*5)
	defer cancel()

	if err := u.UserHandler.Delete(ctx, username); err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return utils.ErrorMessage(c, &utils.ApiNotFound, err.Error())
		}
		return utils.ErrorMessage(c, &utils.ApiInternalServer, err.Error())
	}

	return utils.SuccessMessage(c, &utils.ApiDelete, nil)
}

func (u *userImpl) FindAllUser(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Second*5)
	defer cancel()

	users, err := u.UserHandler.FindAllUser(ctx)
	if err != nil {
		return utils.ErrorMessage(c, &utils.ApiInternalServer, nil)
	}
	if len(users) == 0 {
		return utils.ErrorMessage(c, &utils.ApiNotFound, "empty data")
	}

	return utils.SuccessMessage(c, &utils.ApiOk, users)
}

func (u *userImpl) Update(c echo.Context) error {
	username := c.Param("username")

	payload := web.UpdateReq{}
	if err := c.Bind(&payload); err != nil {
		return utils.ErrorMessage(c, &utils.ApiBadRequest, "invalid request body")
	}

	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Second*5)
	defer cancel()

	if err := u.UserHandler.Update(ctx, username, &payload); err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return utils.ErrorMessage(c, &utils.ApiNotFound, err.Error())
		}
		if strings.Contains(err.Error(), "password constraint") {
			return utils.ErrorMessage(c, &utils.ApiBadRequest, err.Error())
		}
		return utils.ErrorMessage(c, &utils.ApiInternalServer, nil)
	}

	return utils.SuccessMessage(c, &utils.ApiOk, "updated successfuly")
}
