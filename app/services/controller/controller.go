package controller

import "tpk-backend/app/services"

type controllerTPK struct {
	service services.ServiceInterface
}

func NewController(service services.ServiceInterface) services.ControllerInterface {
	return &controllerTPK{
		service: service,
	}
}