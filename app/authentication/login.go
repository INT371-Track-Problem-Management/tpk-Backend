package authentication

import (
	"errors"
	"log"
	"tpk-backend/app/model/entity"
	"tpk-backend/app/model/request"
	"tpk-backend/app/service/repositories"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Login(ctx echo.Context, conn *gorm.DB, req request.User) (*string, error) {
	user, err := GetUser(conn, req)
	if err != nil {
		return nil, err
	}

	if req.Email != user.Email || req.Password != user.Password {
		errUn := errors.New("Unatutherize")
		return nil, errUn
	}

	if user.Role == "C" {
		cus, err := repositories.CustomerByEmail(ctx, conn, user.Email)
		if err != nil {
			return nil, err
		}
		if cus.Status == "I" {
			errorstatus := errors.New(`plese activate your account befor login please check your email`)
			log.Println(errorstatus)
			return nil, errorstatus

		}
	}

	token, err := GenerateTokenLogin(user.Email, user.Role)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return token, nil
}

func GetUser(conn *gorm.DB, req request.User) (*entity.User, error) {
	user := new(entity.User)
	err := conn.Table("userMaster").Where("email = ?", req.Email).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
