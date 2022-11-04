package services

import (
	"tpk-backend/app/models/request"
	"tpk-backend/app/models/response"
)

type ServiceInterface interface {
	CheckHealthy() (*string, error)

	Login(req request.User) (*response.Token, error)
	RegisterCustomersService(req request.RegisterCustomer) (*int, error)
	RegisterOwnerService(req request.RegisterOwner) (*int, error)
}
