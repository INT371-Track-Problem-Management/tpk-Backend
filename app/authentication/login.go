package authentication

import (
	"errors"
	"fmt"
	"log"
	"time"
	"tpk-backend/app/model/entity"
	"tpk-backend/app/model/request"
	"tpk-backend/app/service/repositories"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type JwtCustomClaims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	Status   bool   `json:"status"`
	jwt.StandardClaims
}

type JwtRegisterActivate struct {
	CustomerId int
	jwt.StandardClaims
}

func GenerateTokenRegister(cusId int) (*string, error) {
	claims := &JwtRegisterActivate{
		cusId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte("abcdefghijkmn"))
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func Login(ctx echo.Context, conn *gorm.DB, req request.User) (*string, error) {
	user, err := GetUser(ctx, conn, req)
	if err != nil {
		return nil, err
	}

	if req.Username != user.Username || req.Password != user.Password {
		errUn := errors.New("Unatutherize")
		return nil, errUn
	}

	cus, err := repositories.CustomerByUsername(ctx, conn, user.Username)
	if err != nil {
		return nil, err
	}
	if cus.Status == "I" {
		nonAct := fmt.Sprintln("nonAct")
		errorstatus := errors.New(nonAct)
		log.Println(errorstatus)
		return nil, errorstatus

	}
	// Set custom claims
	claims := &JwtCustomClaims{
		user.Username,
		user.Role,
		true,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte("abcdefghijkmn"))
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func GetUser(ctx echo.Context, conn *gorm.DB, req request.User) (*entity.User, error) {
	user := new(entity.User)
	err := conn.Table("user").Where("username = ?", req.Username).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
