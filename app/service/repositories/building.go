package repositories

import (
	"tpk-backend/app/model/entity"
	"tpk-backend/app/model/request"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Building(ctx echo.Context, conn *gorm.DB, req request.Building) (*entity.Building, error) {
	var building entity.Building
	err := conn.Table("building").Where("buildingId = ?", req.BuildingId).Find(&building).Error
	if err != nil {
		return nil, err
	}
	return &building, nil
}

func BuildingInsert(ctx echo.Context, conn *gorm.DB, model entity.BuildingInsert) error {
	err := conn.Table("building").Create(&model).Error
	if err != nil {
		return err
	}
	return nil
}

func BuildingDelete(ctx echo.Context, conn *gorm.DB, req request.BuildingDelete) error {
	var err error
	err = conn.Exec("DELETE FROM room WHERE buildingId = ?", req.BuildingId).Error
	if err != nil {
		return err
	}
	err = conn.Exec("DELETE FROM building WHERE buildingId = ?", req.BuildingId).Error
	if err != nil {
		return err
	}
	return nil
}
