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

type ResponseToken struct {
	Token string
	Name  string
}

func Login(ctx echo.Context, conn *gorm.DB, req request.User) (*ResponseToken, error) {
	user, err := GetUser(conn, req.Email)
	if err != nil {
		return nil, err
	}

	if req.Email != user.Email {
		errUn := errors.New("Unatutherize")
		return nil, errUn
	}

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
		err = SaveToken(conn, token)
		if err != nil {
			return nil, err
		}
		res := ResponseToken{
			Token: *token,
			Name:  cus.Fname,
		}
		return &res, nil
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
		err = SaveToken(conn, token)
		if err != nil {
			return nil, err
		}

		res := ResponseToken{
			Token: *token,
			Name:  emp.Fname,
		}
		return &res, nil
	}

	if user.Role == "A" {
		emp, err := repositories.EmployeeByEmail(ctx, conn, user.Email)
		if err != nil {
			return nil, err
		}
		token, err := GenerateTokenLogin(emp.EmployeeId, user.Email, user.Role)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		err = SaveToken(conn, token)
		if err != nil {
			return nil, err
		}

		res := ResponseToken{
			Token: *token,
			Name:  emp.Fname,
		}
		return &res, nil
	}

	return nil, nil
}

func GetUser(conn *gorm.DB, email string) (*entity.User, error) {
	user := new(entity.User)
	err := conn.Table("userApp").Where("email = ?", email).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func SaveToken(conn *gorm.DB, token *string) error {
	save := entity.SaveToken{
		Token:  *token,
		Status: `A`,
	}
	err := conn.Table("tokenApp").Create(save).Error
	if err != nil {
		return err
	}
	return nil
}
