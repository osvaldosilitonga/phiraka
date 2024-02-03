package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/osvaldosilitonga/phiraka/server/configs"
)

func Routes(e *echo.Echo) {
	db := configs.InitDB()

	v1 := e.Group("/api/v1")
	{
		v1.GET("", func(c echo.Context) error {
			return c.JSON(200, echo.Map{
				"msg": "OK",
			})
		})
	}
}
