package utils

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/osvaldosilitonga/phiraka/server/domain/web"
)

var (
	ApiBadRequest = web.ErrorResp{
		Code:   http.StatusBadRequest,
		Status: "Bad Request",
	}

	ApiNotFound = web.ErrorResp{
		Code:   http.StatusNotFound,
		Status: "not found",
	}

	ApiForbidden = web.ErrorResp{
		Code:   http.StatusForbidden,
		Status: "forbidden",
	}

	ApiInternalServer = web.ErrorResp{
		Code:   http.StatusInternalServerError,
		Status: "internal server error",
	}

	ApiUnauthorized = web.ErrorResp{
		Code:   http.StatusUnauthorized,
		Status: "unauthorized",
	}
)

func ErrorMessage(c echo.Context, apiErr *web.ErrorResp, message any) error {
	apiErr.Message = message

	return echo.NewHTTPError(apiErr.Code, apiErr)
}
