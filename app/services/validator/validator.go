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

func (v validator) StatusToken(token string) (bool, error) {
	var status string
	if err := v.db.Table("tokenApp").Where("token = ?", token).Select("status").Scan(&status).Error; err != nil {
		return false, err
	}
	if status == "A" {
		return true, nil
	}
	return false, nil
}
