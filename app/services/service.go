package services

import (
	"mime/multipart"
	"tpk-backend/app/models/model"
	"tpk-backend/app/models/request"
	"tpk-backend/app/models/response"
)

type ServiceInterface interface {
	CheckHealthy() (*string, error)
	Login(req request.User) (*response.Token, error)
	LogoutToken(token string) error
	RegisterCustomersService(req request.RegisterCustomer) (*int, error)
	RegisterOwnerService(req request.RegisterOwner) (*int, error)
	CreateReport(req request.ReportInsert, image *multipart.FileHeader) (*int, error)
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
	FetchCustomerByEmail(email string) (*model.Customer, error)
	FetchEmployeeByEmail(email string) (*model.Employee, error)
	RoomByBuildingId(buildingId int) (*response.RoomByBuildingId, error)
	GetAllRoomAndCustomerByBuildingId(buildingId int) ([]*model.RoomJoinBulding, error)
	Maintainerlist() ([]*model.Maintainer, error)
	MaintainerById(maintainerId int) (*model.Maintainer, error)
	EditProfile(req request.EditProfile, email string, role string) error
	GetListCustomer() ([]*model.Customer, error)
	GetListEmployee() ([]*model.Employee, error)
	FetchEmployeeById(customerId int) (*model.Employee, error)
	FetchCustomerById(customerId int) (*model.Customer, error)
	DeleteCustomer(id int) error
	DeleteEmployee(id int) error
	EndJobReport(req request.EndJobReport) error
	FetcStatDashBoard(req request.Stat) (*model.Stat, error)
	FetchStatMaintain() (*model.StatMaintainer, error)
	FetchOverviewMaintain(maintainerId int) (*response.OverviewMaintainer, error)
	ReportMediaById(id string) (*model.ReportMedia, error)
}
