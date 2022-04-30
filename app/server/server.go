package server

import (
	"fmt"
	"net/http"
	"tpk-backend/app/model/request"
	"tpk-backend/app/service/controller"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func StartServer() {
	e := echo.New()
	h := TestHanler{}
	h.Initialize()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/api/test", h.Test)

	e.Logger.Fatal(e.Start(":5000"))
}

type TestHanler struct {
	DB *gorm.DB
}

func (h *TestHanler) Initialize() {

	dns := "dev:123456789@tcp(52.139.153.111:3306)/project?charset=utf8&parseTime=True&loc=Local"
	conn, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println(err)
	}
	h.DB = conn
}

func (h *TestHanler) Test(ctx echo.Context) error {
	req := new(request.Test)
	err := ctx.Bind(&req)
	if err != nil {
		return err
	}
	res, err := controller.TestController(ctx, *req, h.DB)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, res)
}
