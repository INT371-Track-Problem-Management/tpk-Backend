package authentication

import (
	"errors"
	"time"
	"tpk-backend/app/model/entity"
	"tpk-backend/app/model/request"

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

func Login(ctx echo.Context, conn *gorm.DB, req request.User) (*string, error) {
	user, err := GetUser(ctx, conn, req)
	if err != nil {
		return nil, err
	}

	if req.Username != user.Username || req.Password != user.Password {
		errUn := errors.New("Unatutherize")
		return nil, errUn
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
