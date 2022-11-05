package services

import (
	"tpk-backend/app/models/model"
	"tpk-backend/app/models/request"
)

type RepositoryInterface interface {
	CheckHealthy() (*string, error)

	//userApp
	GetUser(email string) (*model.User, error)
	SaveToken(token *string) error
	RegisUser(req model.User) error

	//customer
	CustomerByEmail(email string) (*model.Customer, error)
	GetCustomerById(customerId int) (*model.Customer, error)
	RegisterCustomersRepo(req request.CustomerRegis) (*int, error)

	//employee
	EmployeeById(id int) (*model.Employee, error)
	EmployeeByEmail(email string) (*model.Employee, error)
	RegisterEmployeeRepo(req model.EmployeeRegis) (*int, error)

	//report
	Report() (*[]model.ReportJoinEngage, error)
	ReportByCreatedBy(customerId string) (*[]model.Report, error)
	ReportById(reportId int) (*model.Report, error)
	ReportInsert(req model.ReportInsert) (*int, *model.Customer, error)
	ReportChangeStatus(req model.ReportChangeStatus) error
	DeleteReportById(req request.Report) error
	ReportByRoomId(roomId string) (*[]model.ReportJoinEngage, error)
	ReportListForCustomer(customerId string) (*[]model.ReportJoinEngage, error)

	//reportStatus
	ReportStatusByReportId(reportId string) (*[]model.ReportStatus, error)
}
