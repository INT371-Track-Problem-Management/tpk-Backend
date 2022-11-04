package service

import (
	"tpk-backend/app/model/entity"
	"tpk-backend/app/model/request"
	"tpk-backend/app/pkg"
	"tpk-backend/app/service/repositories"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func AddMaintainer(ctx echo.Context, conn *gorm.DB, req request.Maintainer) (*int, error) {
	now := pkg.GetDatetime()
	entity := entity.AddMaintainer{
		Fname:    req.Fname,
		Lname:    req.Lname,
		Phone:    req.Phone,
		CreateAt: now,
		UpdateAt: now,
		UpdateBy: req.UpdateBy,
	}
	id, err := repositories.AddMaintainer(ctx, conn, entity)
	if err != nil {
		return nil, err
	}
	return id, nil
}

func Maintainerlist(ctx echo.Context, conn *gorm.DB) ([]*entity.Maintainer, error) {
	maintainers, err := repositories.Maintainerlist(ctx, conn)
	if err != nil {
		return nil, err
	}
	return maintainers, nil
}

func MaintainerById(ctx echo.Context, conn *gorm.DB, maintainerId string) (*entity.Maintainer, error) {
	maintainer, err := repositories.MaintainerById(ctx, conn, maintainerId)
	if err != nil {
		return nil, err
	}
	return maintainer, nil
}
