package repository

import (
	"tpk-backend/app/services"

	"gorm.io/gorm"
)

type mysqlRepository struct {
	conn *gorm.DB
}

func NewRepository(conn *gorm.DB) services.RepositoryInterface {
	return &mysqlRepository{
		conn: conn,
	}
}
