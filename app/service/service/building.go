package service

import (
	"fmt"
	"tpk-backend/app/model/entity"
	"tpk-backend/app/model/request"
	"tpk-backend/app/model/response"
	"tpk-backend/app/pkg"
	"tpk-backend/app/service/repositories"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Building(ctx echo.Context, conn *gorm.DB, req request.Building) (*response.Building, error) {
	data, err := repositories.Building(ctx, conn, req)
	if err != nil {
		return nil, err
	}
	fmt.Println(data)
	res := &response.Building{
		BuildingId:  data.BuildingId,
		Address: data.Address,
		Phone:   data.Phone,
		Email:   data.Email,
	}
	return res, nil
}

func BuildingInsert(ctx echo.Context, conn *gorm.DB, req request.BuildingInsert) (string, error) {
	timenow := pkg.GetDatetime()
	model := entity.BuildingInsert{
		BuildingName: req.BuildingName,
		CreateAt:     timenow,
		UpdateAt:     timenow,
		UpdateBy:     req.UpdateBy,
	}
	err := repositories.BuildingInsert(ctx, conn, model)
	if err != nil {
		return "Can not insert building", err
	}
	return "Insert success", nil
}

func BuildingDelete(ctx echo.Context, conn *gorm.DB, req request.BuildingDelete) (string, error) {
	err := repositories.BuildingDelete(ctx, conn, req)
	if err != nil {
		return "Can not delete", err
	}
	return "Delete success", nil
}
