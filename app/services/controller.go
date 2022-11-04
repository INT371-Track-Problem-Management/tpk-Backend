package services

import "github.com/labstack/echo/v4"

type ControllerInterface interface {
	CheckHealthy(ctx echo.Context) error
}
