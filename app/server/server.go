package server

import (
	"fmt"
	"net/http"
	"tpk-backend/app/model/entity"
	"tpk-backend/app/model/request"
	"tpk-backend/app/model/response"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func StartServer() {
	e := echo.New()
	h := TestHanler{}
	h.Initialize()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/test", h.Test)

	e.Logger.Fatal(e.Start(":5000"))
}

type TestHanler struct {
	DB *gorm.DB
}

func (h *TestHanler) Initialize() {
	dns := "dev:123456789@tcp(52.139.153.111:3306)/project?charset=utf8&parseTime=True&loc=Local"
	conn, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	h.DB = conn
}

func (h *TestHanler) Test(ctx echo.Context) error {
	userid := ctx.QueryParam("userId")
	data := new(entity.Test)
	req := new(request.Test)
	err := ctx.Bind(&req)
	if err != nil {
		return err
	}
	err = h.DB.Table("test").Select("USERID", "NAME").Where("USERID = ?", userid).Find(&data).Error
	if err != nil {
		return err
	}

	res := &response.Test{
		UserId: data.UserId,
		Name:   data.Name,
	}

	return ctx.JSON(http.StatusOK, res)
}
