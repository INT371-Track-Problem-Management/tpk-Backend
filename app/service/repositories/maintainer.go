package repositories

import (
	"tpk-backend/app/model/entity"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func AddMaintainer(ctx echo.Context, conn *gorm.DB, req entity.AddMaintainer) (*int, error) {
	stmt := conn.Begin()
	var err error
	err = stmt.Table("maintainer").Create(&req).Error
	if err != nil {
		return nil, err
	}
	var id int
	err = stmt.Table("maintainer").Select("maintainerId").Where("fname = ?", req.Fname).Where("lname = ?", req.Lname).Scan(&id).Error
	if err != nil {
		return nil, err
	}
	stmt.Commit()
	return &id, nil
}

func Maintainerlist(ctx echo.Context, conn *gorm.DB) ([]*entity.Maintainer, error) {
	var maintainers []*entity.Maintainer
	err := conn.Table("maintainer").Find(&maintainers).Error
	if err != nil {
		return nil, err
	}
	return maintainers, nil
}

func MaintainerById(ctx echo.Context, conn *gorm.DB, maintainerId string) (*entity.Maintainer, error) {
	var maintainers entity.Maintainer
	err := conn.Table("maintainer").Where("maintainerId = ?", maintainerId).Scan(&maintainers).Error
	if err != nil {
		return nil, err
	}
	return &maintainers, nil
}
