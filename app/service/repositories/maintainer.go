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
