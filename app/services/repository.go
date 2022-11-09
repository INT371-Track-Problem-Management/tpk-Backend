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

	//customer
	CustomerByEmail(email string) (*model.Customer, error)
	GetCustomerById(customerId int) (*model.Customer, error)
	RegisterCustomersRepo(req request.CustomerRegis) (*int, error)
	GetEmailCreateByReportId(reportId int) (*string, error)

	//employee
	EmployeeById(id int) (*model.Employee, error)
	EmployeeByEmail(email string) (*model.Employee, error)
	RegisterEmployeeRepo(req model.EmployeeRegis) (*int, error)

	//report
	CreateReport(model model.ReportInsert, session *gorm.DB) (*int, error)
	ListReport(fillter *request.FillterReport) (*[]response.ReportList, error)
	ReportDetailById(reportId int) (*model.Report, error)
	ChangeCategory(req request.ReportChangeCategory) error
	ChangeStatusReport(req request.ReportStatus, session *gorm.DB) error

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

	//bulding
	AllBuilding() (*[]response.AllBuilding, error)
	CreateBuilding(model model.BuildingInsert) (*int64, error)
	BuildingDelete(buildingId int) error

	//maintainer
	CreateMaintainer(req model.AddMaintainer) error
}
