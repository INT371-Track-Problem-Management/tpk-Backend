package repositories

import (
	"tpk-backend/app/model/entity"
	"tpk-backend/app/model/request"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Dorm(ctx echo.Context, conn *gorm.DB, req request.Dorm) (*entity.Dorm, error) {
	var dorm entity.Dorm
	err := conn.Table("dorm").Where("dormId = ?", req.DormId).Find(&dorm).Error
	if err != nil {
		return nil, err
	}
	return &dorm, nil
}
