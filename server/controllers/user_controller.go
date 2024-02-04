package controllers

import (
	"context"
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
