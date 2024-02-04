package utils

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/osvaldosilitonga/phiraka/server/domain/web"
)

var (
	ApiOk = web.SuccessResp{
		Code:   http.StatusOK,
		Status: "ok",
	}

	ApiDelete = web.SuccessResp{
		Code:   http.StatusOK,
		Status: "delete success",
	}

	ApiCreate = web.SuccessResp{
		Code:   http.StatusCreated,
		Status: "create success",
	}

	ApiUpdate = web.SuccessResp{
		Code:   http.StatusOK,
		Status: "update success",
	}
)

func SuccessMessage(c echo.Context, apiSuccess *web.SuccessResp, data any) error {
	apiSuccess.Data = data

	return c.JSON(apiSuccess.Code, apiSuccess)
}
