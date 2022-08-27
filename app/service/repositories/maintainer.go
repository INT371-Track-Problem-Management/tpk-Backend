package repositories

import (
	"tpk-backend/app/model/request"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func AddMaintainer(ctx echo.Context, conn *gorm.DB, req request.Maintainer) (*int, error) {
	var err error
	err = conn.Table("maintainer").Create(&req).Error
	if err != nil {
		return nil, err
	}
	var id int
	err = conn.Table("maintainer").Select("maintainerId").Where("fnamme = ?", req.Fname).Where("lname = ?", req.Lname).Scan(&id).Error
	if err != nil {
		return nil, err
	}
	return &id, nil
}
