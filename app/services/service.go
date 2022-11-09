package services

import (
	"tpk-backend/app/models/model"
	"tpk-backend/app/models/request"
	"tpk-backend/app/models/response"
)

type ServiceInterface interface {
	CheckHealthy() (*string, error)
	Login(req request.User) (*response.Token, error)
	RegisterCustomersService(req request.RegisterCustomer) (*int, error)
	RegisterOwnerService(req request.RegisterOwner) (*int, error)
	CreateReport(req request.ReportInsert) (*int, error)
	ListReport(fillter *request.FillterReport) (*[]response.ReportList, error)
	ReportDetailById(reportId int) (*model.Report, error)
	ReportStatusByReportId(reportId string) (*[]model.ReportStatus, error)
	ReportEnagegeFixDateDetail(reportId string) (*response.ReportEngageFixDate, error)
	SelectedPlanFixDate(req request.SelectedPlanFixDate) error
	ChangeCategory(req request.ReportChangeCategory) error
	NewFixDate(req request.ReportEngage) error
	ChangeStatusReport(req request.ReportChangeStatus) error
	GetRoomWithCustomerId(customerId int) (*[]model.RoomWithCustomerId, error)
	AllBuilding() (*[]response.AllBuilding, error)
	CreateBuilding(req request.BuildingInsert) (*int64, error)
	GetAllRoomWithCustomerByBuildingId(buildingId int) ([]*model.RoomJoinBulding, error)
	CreateRoom(req request.RoomInsert) error
	BuildingDelete(buildingId int) error
	ChangeEmail(req request.ChangeEmail, oldEmail string) error
	ChangePassword(req request.ChangePassword) error
	AssignJobMaintainer(req request.AssignReport) error
	CreateMaintainer(req request.Maintainer) error
	AddCustomerIntoRoom(req request.RoomAddCustomer) error
	RemoveCustomerFromRoom(id int) error
}
