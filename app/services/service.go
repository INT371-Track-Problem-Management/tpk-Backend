package services

import (
	"tpk-backend/app/models/request"
	"tpk-backend/app/models/response"
)

type ServiceInterface interface {
	CheckHealthy() (*string, error)

	Login(req request.User) (*response.Token, error)
}
