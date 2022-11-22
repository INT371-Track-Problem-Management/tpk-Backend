package services

import "github.com/labstack/echo/v4"

type ValidatorInterface interface {
	AdminValidation(ctx echo.HandlerFunc) echo.HandlerFunc
	CustomerValidation(ctx echo.HandlerFunc) echo.HandlerFunc
	EmployeeValidation(ctx echo.HandlerFunc) echo.HandlerFunc
	StatusToken(token string) bool
}
