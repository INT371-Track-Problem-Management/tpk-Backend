package services

import "tpk-backend/app/models/model"

type RepositoryInterface interface {
	CheckHealthy() (*string, error)

	//userApp
	GetUser(email string) (*model.User, error)
	SaveToken(token *string) error

	//customer
	CustomerByEmail(email string) (*model.Customer, error)
	GetCustomerById(customerId int) (*model.Customer, error)

	//employee
	EmployeeById(id int) (*model.Employee, error)
	EmployeeByEmail(email string) (*model.Employee, error)
}
