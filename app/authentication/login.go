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

	if req.Email != user.Email {
		errUn := errors.New("Unatutherize")
		return nil, errUn
	}

	log.Println(req.Password)
	log.Println(user.Password)

	if pwd := ComparePassword(req.Password, user.Password); !pwd {
		errUn := errors.New("Unatutherize")
		return nil, errUn
	}

	if user.Role == "C" {
		cus, err := repositories.CustomerByEmail(ctx, conn, user.Email)
		if err != nil {
			return nil, err
		}
		token, err := GenerateTokenLogin(cus.CustomerId, user.Email, user.Role)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		return token, nil
	}

	if user.Role == "E" {
		emp, err := repositories.EmployeeByEmail(ctx, conn, user.Email)
		if err != nil {
			return nil, err
		}
		token, err := GenerateTokenLogin(emp.EmployeeId, user.Email, user.Role)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		return token, nil
	}

	return nil, nil
}

func GetUser(conn *gorm.DB, req request.User) (*entity.User, error) {
	user := new(entity.User)
	err := conn.Table("userMaster").Where("email = ?", req.Email).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
