package services

import (
	"tpk-backend/app/models/model"
	"tpk-backend/app/models/request"
	"tpk-backend/app/models/response"

	"gorm.io/gorm"
)

type RepositoryInterface interface {
	CheckHealthy() (*string, error)

	//userApp
	GetUser(email string) (*model.User, error)
	SaveToken(token *string, role string) error
	RegisUser(req model.User) error
	ChangeEmail(req request.ChangeEmail, oldEmail string) error
	ChangePassword(model model.ChangePassword) error
	EditProfile(model model.EditProfile, email string, role string) error
	LogoutToken(token string) error

	//customer
	CustomerByEmail(email string) (*model.Customer, error)
	GetCustomerById(customerId int) (*model.Customer, error)
	RegisterCustomersRepo(req request.CustomerRegis) (*int, error)
	GetEmailCreateByReportId(reportId int) (*string, error)
	GetListCustomer() ([]*model.Customer, error)
	DeleteCustomer(id int) error

	//employee
	EmployeeById(id int) (*model.Employee, error)
	EmployeeByEmail(email string) (*model.Employee, error)
	RegisterEmployeeRepo(req model.EmployeeRegis) (*int, error)
	GetListEmployee() ([]*model.Employee, error)
	DeleteEmployee(id int) error

	//report
	CreateReport(model model.ReportInsert, session *gorm.DB) (*int, error)
	ListReport(fillter *request.FillterReport) (*[]response.ReportList, error)
	ReportDetailById(reportId int) (*model.Report, error)
	ChangeCategory(req request.ReportChangeCategory) error
	ChangeStatusReport(req request.ReportStatus, session *gorm.DB) error
	EndJobReport(session *gorm.DB, req model.EndJobReport) error
	FetcStatDashBoard(req request.Stat) (*model.Stat, error)

	//reportStatus
	CreateReportStatus(model request.ReportStatus, session *gorm.DB) error
	ReportStatusByReportId(reportId string) (*[]model.ReportStatus, error)

	//reportEngage
	CreateReporEngage(model model.InsertReportEngage, session *gorm.DB) (*int, error)
	ReportEnagegeByReportId(reportId string) (*model.ReportEngage, error)
	SelectedPlanFixDate(req request.SelectedPlanFixDate) error
	EditEngage(req request.ReportEngage, session *gorm.DB) error
	AssignJobMaintainer(req request.AssignReport) error

	//fixdate
	CreateFixdate(model model.CreateFixdate, session *gorm.DB) error
	FixdateByEngageId(engage model.ReportEngage) (*[]model.Fixdate, error)

	//room
	GetRoomWithCustomerId(customerId int) (*[]model.RoomWithCustomerId, error)
	GetAllRoomWithCustomerByBuildingId(buildingId int) ([]*model.RoomJoinBulding, error)
	CreateRoom(model model.RoomInsert, session *gorm.DB) error
	RoomByRoomId(roomId int) (*model.Room, error)
	TotalFlooorsByBuildingId(buildingId int) (*int, error)
	RoomInFloorByBuildingId(buildingId int, floor int) (*[]model.RoomByFloors, error)
	GetAllRoomAndCustomerByBuildingId(buildingId int) ([]*model.RoomJoinBulding, error)

	//bulding
	AllBuilding() (*[]response.AllBuilding, error)
	CreateBuilding(model model.BuildingInsert) (*int64, error)
	BuildingDelete(buildingId int) error

	//maintainer
	CreateMaintainer(req model.AddMaintainer) error
	Maintainerlist() ([]*model.Maintainer, error)
	MaintainerById(maintainerId int) (*model.Maintainer, error)
	FetchStatMaintain() (*model.StatMaintainer, error)
	FetchOverviewMaintain(id int) (*[]model.OverviewMaintainer, error)

	//RoomWithCustomer
	AddCustomerIntoRoom(model model.RoomAddCustomer) error
	RemoveCustomerFromRoom(id int, now string) error
}
