package service

import (
	"tpk-backend/app/services"

	"gorm.io/gorm"
)

type serviceTPK struct {
	repo     services.RepositoryInterface
	database *gorm.DB
}

func NewService(conn services.RepositoryInterface, db *gorm.DB) services.ServiceInterface {
	return &serviceTPK{
		repo:     conn,
		database: db,
	}
}
