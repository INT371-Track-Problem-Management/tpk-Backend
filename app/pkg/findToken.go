package pkg

import (
	"tpk-backend/app/model/entity"

	"gorm.io/gorm"
)

var conn *gorm.DB

func FindToken(token string) (*entity.SaveToken, error) {
	var t *entity.SaveToken
	err := conn.Table("customer").Where("token = ?", token).Find(&t).Error
	if err != nil {
		return nil, err
	}
	return t, nil
}
