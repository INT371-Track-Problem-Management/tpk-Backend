package repositories

import (
	"tpk-backend/app/model/entity"
	"tpk-backend/app/model/request"
	"tpk-backend/app/model/response"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func AllBuilding(ctx echo.Context, conn *gorm.DB) (*[]response.AllBuilding, error) {
	var building []response.AllBuilding
	err := conn.Table("building").Find(&building).Error
	if err != nil {
		return nil, err
	}
	return &building, nil
}

func BuildingById(ctx echo.Context, conn *gorm.DB, buildingId string) (*entity.Building, error) {
	var building entity.Building
	err := conn.Table("building").Where("buildingId = ?", buildingId).Find(&building).Error
	if err != nil {
		return nil, err
	}
	return &building, nil
}

func BuildingInsert(ctx echo.Context, conn *gorm.DB, model entity.BuildingInsert) (*int64, error) {
	err := conn.Table("building").Create(&model).Error
	if err != nil {
		return nil, err
	}
	var id int64
	err = conn.Table("building").Where("createAt = ?", model.CreateAt).Select("buildingId").Scan(&id).Error
	if err != nil {
		return nil, err
	}
	return &id, nil
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
