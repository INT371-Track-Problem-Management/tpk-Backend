package repositories

import (
	"tpk-backend/app/model/entity"
	"tpk-backend/app/model/request"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisUser(ctx echo.Context, conn *gorm.DB, req entity.User) error {
	err := conn.Table("userApp").Create(&req).Error
	if err != nil {
		return err
	}
	return nil
}

func ChangeEmail(ctx echo.Context, conn *gorm.DB, req request.ChangeEmail, oldEmail string) error {
	var err error
	stmt := conn.Begin()
	err = stmt.Table("userApp").Where("email = ?", oldEmail).Update("email = ?", req.NewEmail).Error
	if err != nil {
		stmt.Rollback()
		return err
	}
	stmt.Commit()
	return nil
}
