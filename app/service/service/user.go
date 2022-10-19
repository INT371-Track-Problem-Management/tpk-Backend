package service

import (
	"fmt"
	"net/http"
	"tpk-backend/app/authentication"
	"tpk-backend/app/model/entity"
	"tpk-backend/app/model/request"
	"tpk-backend/app/service/repositories"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ChangeEmail(ctx echo.Context, conn *gorm.DB, req request.ChangeEmail, oldEmail string) error {
	var err error
	user := request.User{
		Email: oldEmail,
	}

	oldpwd, err := authentication.GetUser(conn, user)
	if err != nil {
		return err
	}

	if req.Password != oldpwd.Password {
		return ctx.JSON(http.StatusBadRequest, "Invalid Token")
	}

	err = repositories.ChangeEmail(ctx, conn, req, oldEmail)
	if err != nil {
		return err
	}
	return nil
}

func ChangePassword(ctx echo.Context, conn *gorm.DB, req request.ChangePassword) error {
	model := entity.ChangePassword{
		Email:       req.Email,
		OldPassword: req.OldPassword,
		NewPassword: req.NewPassword,
	}
	fmt.Println(model.NewPassword)
	err := repositories.ChangePassword(ctx, conn, model)
	if err != nil {
		return err
	}
	return nil
}
