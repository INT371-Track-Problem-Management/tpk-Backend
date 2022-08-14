package controller

import (
	"tpk-backend/app/model/request"
	"tpk-backend/app/service/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ChangeEmail(ctx echo.Context, conn *gorm.DB, req request.ChangeEmail, oldEmail string) error {
	err := service.ChangeEmail(ctx, conn, req, oldEmail)
	if err != nil {
		return err
	}
	return nil
}
