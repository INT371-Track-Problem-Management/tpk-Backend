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
	ChangeCategory(ctx echo.Context) error
	NewFixDate(ctx echo.Context) error
	ChangeStatusReport(ctx echo.Context) error
	GetRoomWithCustomerId(ctx echo.Context) error
	AllBuilding(ctx echo.Context) error
	CreateBuilding(ctx echo.Context) error
	CreateRoom(ctx echo.Context) error
	BuildingDelete(ctx echo.Context) error
	ChangeEmail(ctx echo.Context) error
	ChangePassword(ctx echo.Context) error
	AssignJobMaintainer(ctx echo.Context) error
	CreateMaintainer(ctx echo.Context) error
	AddCustomerIntoRoom(ctx echo.Context) error
	RemoveCustomerFromRoom(ctx echo.Context) error
	FetchProfile(ctx echo.Context) error
}
