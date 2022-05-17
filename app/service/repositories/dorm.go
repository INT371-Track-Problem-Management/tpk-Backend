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

func DormInsert(ctx echo.Context, conn *gorm.DB, req request.DormInsert) error {
	err := conn.Table("dorm").Create(&req).Error
	if err != nil {
		return err
	}
	return nil
}

func DormDelete(ctx echo.Context, conn *gorm.DB, req request.DormDelete) error {
	var err error
	err = conn.Exec("DELETE FROM room WHERE dormId = ?", req.DormId).Error
	if err != nil {
		return err
	}
	err = conn.Exec("DELETE FROM dorm WHERE dormId = ?", req.DormId).Error
	if err != nil {
		return err
	}
	return nil
}
