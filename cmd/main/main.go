package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/aimerzarashi/oqs/internal/ui/hello"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	hello.RegisterHandlers(e)
	
	e.Logger.Fatal(e.Start(":1323"))
}
