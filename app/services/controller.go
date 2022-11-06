package services

import (
	"github.com/labstack/echo/v4"
)

type ControllerInterface interface {
	CheckHealthy(ctx echo.Context) error
	Login(ctx echo.Context) error
	RegisterCustomers(ctx echo.Context) error
	RegisterOwner(ctx echo.Context) error
	CreateReport(ctx echo.Context) error
	ListReport(ctx echo.Context) error
	ReportDetailById(ctx echo.Context) error
	ReportStatusByReportId(ctx echo.Context) error
	ReportEnagegeFixDateDetail(ctx echo.Context) error
	SelectedPlanFixDate(ctx echo.Context) error
}
