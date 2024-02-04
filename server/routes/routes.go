package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/osvaldosilitonga/phiraka/server/configs"
	"github.com/osvaldosilitonga/phiraka/server/controllers"
	"github.com/osvaldosilitonga/phiraka/server/handlers"
	"github.com/osvaldosilitonga/phiraka/server/repositories"
)

func Routes(e *echo.Echo) {
	db := configs.InitDB()

	// Repositories
	userRepository := repositories.NewUserRepository(db)

	// Handler
	userHandler := handlers.NewUserHandler(userRepository)

	// Controller
	userController := controllers.NewUserController(userHandler)

	v1 := e.Group("/api/v1")
	{
		v1.GET("", func(c echo.Context) error {
			return c.JSON(200, echo.Map{
				"msg": "OK",
			})
		})

		v1.POST("/register", userController.Register)
		v1.POST("/login", userController.Login)
	}
}
