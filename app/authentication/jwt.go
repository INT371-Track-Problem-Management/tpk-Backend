package authentication

import (
	"strings"
	"time"
	"tpk-backend/app/pkg/config"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type JwtCustomClaims struct {
	Email  string  `json:"email"`
	Role   string  `json:"role"`
	Status bool    `json:"status"`
	Expire float64 `json:"expire"`
	jwt.StandardClaims
}

type JwtRegisterActivate struct {
	CustomerId int
	jwt.StandardClaims
}

var secrete = config.LoadJWTConfig().Secret
var signingKey = []byte(secrete)

func GenerateTokenRegister(cusId int) (*string, error) {
	claims := &JwtRegisterActivate{
		cusId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(signingKey)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func GenerateTokenLogin(email string, role string) (*string, error) {
	// Set custom claims
	claims := &JwtCustomClaims{
		email,
		role,
		true,
		float64(time.Now().Add(time.Hour * 72).Unix()),
		jwt.StandardClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(signingKey)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func ValidateTokenJWTConfig() middleware.JWTConfig {
	config := middleware.JWTConfig{
		Claims:     &JwtCustomClaims{},
		SigningKey: signingKey,
	}
	return config
}

func ValidateCustomerService(ctx echo.Context) (string, bool) {
	jwt := DecodeJWT(ctx)
	if jwt.Expire < float64(time.Now().Unix()) {
		return "Token is expired", false
	}
	if jwt.Role != "C" && jwt.Status == false {
		return "Token can't use", false
	}
	return "Token can use", true
}

func ValidateOwnerService(ctx echo.Context) (string, bool) {
	jwt := DecodeJWT(ctx)
	if jwt.Expire < float64(time.Now().Unix()) {
		return "Token is expired", false
	}
	if jwt.Role != "E" || jwt.Status == false {
		return "Token can't use", false
	}
	return "Token can use", true
}

func DecodeJWT(ctx echo.Context) JwtCustomClaims {
	reqToken := ctx.Request().Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]
	claimsDecode := jwt.MapClaims{}
	token, _ := jwt.ParseWithClaims(reqToken, claimsDecode, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	claims := token.Claims.(jwt.MapClaims)
	email := claims["email"].(string)
	role := claims["role"].(string)
	status := claims["status"].(bool)
	expire := claims["expire"].(float64)
	jwtDecode := JwtCustomClaims{
		Email:          email,
		Role:           role,
		Status:         status,
		Expire:         expire,
		StandardClaims: jwt.StandardClaims{},
	}
	return jwtDecode
}
