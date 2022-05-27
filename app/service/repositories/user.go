package repositories

import (
	"tpk-backend/app/model/entity"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisUser(ctx echo.Context, conn *gorm.DB, req entity.User) error {
	err := conn.Table("user").Create(&req).Error
	if err != nil {
		return err
	}
	return nil
}
