package service

import (
	"tpk-backend/app/model/entity"
	"tpk-backend/app/model/request"
	"tpk-backend/app/model/response"
	"tpk-backend/app/pkg"
	"tpk-backend/app/service/repositories"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func BuildingById(ctx echo.Context, conn *gorm.DB, buildingId string) (*response.Building, error) {
	data, err := repositories.BuildingById(ctx, conn, buildingId)
	if err != nil {
		return nil, err
	}
	res := &response.Building{
		BuildingId:   data.BuildingId,
		BuildingName: data.BuildingName,
		Address:      data.Address,
		Phone:        data.Phone,
		Email:        data.Email,
	}
	return res, nil
}

func AllBuilding(ctx echo.Context, conn *gorm.DB) (*[]response.AllBuilding, error) {
	data, err := repositories.AllBuilding(ctx, conn)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func BuildingInsert(ctx echo.Context, conn *gorm.DB, req request.BuildingInsert) (*int64, error) {
	timenow := pkg.GetDatetime()
	model := entity.BuildingInsert{
		BuildingName: req.BuildingName,
		CreateAt:     timenow,
		UpdateAt:     timenow,
		UpdateBy:     req.UpdateBy,
	}
	id, err := repositories.BuildingInsert(ctx, conn, model)
	if err != nil {
		return nil, err
	}
	return id, nil
}

func BuildingDelete(ctx echo.Context, conn *gorm.DB, req request.BuildingDelete) (string, error) {
	err := repositories.BuildingDelete(ctx, conn, req)
	if err != nil {
		return "Can not delete", err
	}
	return "Delete success", nil
}
