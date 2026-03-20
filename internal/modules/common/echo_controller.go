package common

import "github.com/labstack/echo/v4"

type EchoController interface {
	Register(ec *echo.Group)
}
