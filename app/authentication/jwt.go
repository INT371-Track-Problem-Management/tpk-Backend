package authentication

import (
	"log"
	"strings"
	"time"
	"tpk-backend/app/pkg/config"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type JwtCustomClaims struct {
	Id     int     `json:"id"`
	Email  string  `json:"email"`
	Role   string  `json:"role"`
	Status bool    `json:"status"`
	Expire float64 `json:"expire"`
	jwt.StandardClaims
}

type CheckCustomerApplication struct {
	Id     int
	Token  string
	Status bool
}

type CheckOwnerApplication struct {
	Id     int
	Token  string
	Status bool
}

type JwtRegisterActivate struct {
	CustomerId int
	jwt.StandardClaims
}

type JWTPassword struct {
	Password string
	jwt.StandardClaims
}

var secrete = config.LoadJWTConfig().Secret
var signingKey = []byte(secrete)

func GenerateTokenFromPassword(password string) (*string, error) {
	claims := &JWTPassword{
		password,
		jwt.StandardClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(signingKey)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

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

func GenerateTokenLogin(id int, email string, role string) (*string, error) {
	// Set custom claims
	claims := &JwtCustomClaims{
		id,
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

func ValidateCustomerService(ctx echo.Context) (*CheckCustomerApplication, bool) {
	jwt := DecodeJWT(ctx)
	app := new(CheckCustomerApplication)
	if jwt.Expire < float64(time.Now().Unix()) {
		app.Id = jwt.Id
		app.Token = "Token is expired"
		app.Status = false
		return app, false
	}
	if jwt.Role != "C" && jwt.Status == false {
		app.Id = jwt.Id
		app.Token = "Token can't use"
		app.Status = false
		return app, false
	}
	app.Id = jwt.Id
	app.Token = "Token can use"
	app.Status = true
	return app, true
}

func ValidateOwnerService(ctx echo.Context) (*CheckOwnerApplication, bool) {
	jwt := DecodeJWT(ctx)
	app := new(CheckOwnerApplication)
	if jwt.Expire < float64(time.Now().Unix()) {
		app.Id = jwt.Id
		app.Token = "Token is expired"
		app.Status = false
		return app, false
	}
	if jwt.Role != "E" || jwt.Status == false {
		app.Id = jwt.Id
		app.Token = "Token can't use"
		app.Status = false
		return app, false
	}
	app.Id = jwt.Id
	app.Token = "Token can use"
	app.Status = true
	return app, true
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
	id := claims["id"].(float64)
	email := claims["email"].(string)
	role := claims["role"].(string)
	status := claims["status"].(bool)
	expire := claims["expire"].(float64)
	jwtDecode := JwtCustomClaims{
		Id:             int(id),
		Email:          email,
		Role:           role,
		Status:         status,
		Expire:         expire,
		StandardClaims: jwt.StandardClaims{},
	}
	return jwtDecode
}

func GetTokenFromHeadler(ctx echo.Context) string {
	reqToken := ctx.Request().Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]
	return reqToken
}

func DecodeJWTPassword(password string) JWTPassword {
	claimsDecode := jwt.MapClaims{}
	token, _ := jwt.ParseWithClaims(password, claimsDecode, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	claims := token.Claims.(jwt.MapClaims)
	pwd := claims["Password"].(string)
	jwtDecode := JWTPassword{
		Password:       pwd,
		StandardClaims: jwt.StandardClaims{},
	}
	log.Println(jwtDecode.Password)
	return jwtDecode
}

func ComparePassword(reqPassword string, encrypPWD string) bool {
	pwdDB := DecodeJWTPassword(encrypPWD)
	log.Println(pwdDB)
	return pwdDB.Password == reqPassword
}
