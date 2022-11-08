package validator

import (
	"tpk-backend/app/services"

	"gorm.io/gorm"
)

type validator struct {
	db *gorm.DB
}

func Newvalidator(db *gorm.DB) services.ValidatorInterface {
	return &validator{
		db: db,
	}
}
