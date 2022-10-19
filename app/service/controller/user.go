package controller

import (
	"errors"
	"tpk-backend/app/authentication"
	"tpk-backend/app/constants"
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

func ChangePassword(ctx echo.Context, conn *gorm.DB, req request.ChangePassword) error {
	var err error
	user := request.User{
		Email:    req.Email,
		Password: req.OldPassword,
	}
	getUser, err := authentication.GetUser(conn, user)
	if err != nil {
		return err
	}
	check := authentication.ComparePassword(req.OldPassword, getUser.Password)
	if check == constants.CHECK_FALSE {
		return errors.New("password is not correct")
	}
	encryp, err := authentication.GenerateTokenFromPassword(req.NewPassword)
	if err != nil {
		return err
	}
	req.NewPassword = *encryp
	err = service.ChangePassword(ctx, conn, req)
	if err != nil {
		return err
	}
	return nil
}
