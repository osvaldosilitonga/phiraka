package main

import (
	"fmt"
	"os"

	"github.com/osvaldosilitonga/phiraka/server/initializers"
	"github.com/osvaldosilitonga/phiraka/server/middlewares"
	"github.com/osvaldosilitonga/phiraka/server/routes"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	initializers.LoadEnv()
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.RequestLoggerWithConfig(middlewares.LogrusConfig()))
	e.Validator = &initializers.CustomValidator{Validator: validator.New()}

	routes.Routes(e)

	port := os.Getenv("PORT")
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", port)))
}
