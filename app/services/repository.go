package services

import (
	"tpk-backend/app/models/model"
	"tpk-backend/app/models/request"

	"gorm.io/gorm"
)

type RepositoryInterface interface {
	CheckHealthy() (*string, error)

	//userApp
	GetUser(email string) (*model.User, error)
	SaveToken(token *string, role string) error
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
	CreateReport(model model.ReportInsert, session *gorm.DB) (*int, error)

	//reportStatus
	CreateReportStatus(model request.ReportStatus, session *gorm.DB) error

	//reportEngage
	CreateReporEngage(model model.InsertReportEngage, session *gorm.DB) (*int, error)

	//fixdate
	CreateFixdate(model model.CreateFixdate, session *gorm.DB) error
}
