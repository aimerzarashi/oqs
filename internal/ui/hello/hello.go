package hello

import (
	oapicodegen "github.com/aimerzarashi/oqs/internal/infra/oapicodegen/hello"

	"github.com/labstack/echo/v4"
)

type Api struct {
	oapicodegen.ServerInterface
}

func RegisterHandlers(e *echo.Echo) {
	oapicodegen.RegisterHandlers(e, &Api{})
}
