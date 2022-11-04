package service

import "tpk-backend/app/services"

type serviceTPK struct {
	repo services.RepositoryInterface
}

func NewService(conn services.RepositoryInterface) services.ServiceInterface {
	return &serviceTPK{
		repo: conn,
	}
}
