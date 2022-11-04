package services

import (
	"github.com/labstack/echo/v4"
)

type ControllerInterface interface {
	CheckHealthy(ctx echo.Context) error

	//user
	Login(ctx echo.Context) error
	RegisterCustomers(ctx echo.Context) error
	RegisterOwner(ctx echo.Context) error
}
