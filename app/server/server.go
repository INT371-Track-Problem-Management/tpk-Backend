package server

import (
	"fmt"
	"net/http"
	"os"
	"tpk-backend/app/model/request"
	"tpk-backend/app/service/controller"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func StartServer() {
	key := os.Getenv("KEY")
	port := SetEnv(key)
	fmt.Println("PROJECT RUN ON PORT: " + port)
	e := echo.New()
	h := FuncHandler{}
	h.Initialize()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	api := e.Group("/api/")
	api.GET("test", h.Test)
	api.GET("checkHealthy", h.CheckHealthy)
	api.GET("rooms", h.Rooms)
	api.GET("customer", h.Customer)
	api.PUT("rooms", h.RoomsStatus)
	api.GET("dorm", h.Dorm)
	api.GET("report", h.Report)
	api.GET("reportById", h.ReportById)
	e.Logger.Fatal(e.Start(":" + port))
}

func SetEnv(key string) string {
	var port string
	if key == "PRD" {
		port = "5000"
		return port
	}
	if key == "DEV" {
		port = "3000"
		return port
	} else {
		fmt.Printf("Invalid ENV")
	}
	return ""
}

type FuncHandler struct {
	DB *gorm.DB
}

func (h *FuncHandler) Initialize() {
	dns := "dev:123456789@tcp(52.139.153.111:3306)/project?charset=utf8&parseTime=True&loc=Local"
	conn, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println(err)
	}
	h.DB = conn
}

func (h *FuncHandler) Test(ctx echo.Context) error {
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

func (h *FuncHandler) CheckHealthy(ctx echo.Context) error {
	res, err := controller.CheckHealthy(ctx, h.DB)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) Rooms(ctx echo.Context) error {
	res, err := controller.Rooms(ctx, h.DB)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) RoomsStatus(ctx echo.Context) error {
	req := new(request.RoomsStatus)
	err := ctx.Bind(&req)
	if err != nil {
		return err
	}
	res, err := controller.RoomsStatus(ctx, h.DB, *req)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) Customer(ctx echo.Context) error {
	res, err := controller.Customer(ctx, h.DB)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) Dorm(ctx echo.Context) error {
	req := new(request.Dorm)
	err := ctx.Bind(&req)
	if err != nil {
		return err
	}
	res, err := controller.Dorm(ctx, h.DB, *req)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) Report(ctx echo.Context) error {
	res, err := controller.Report(ctx, h.DB)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) ReportById(ctx echo.Context) error {
	req := new(request.Report)
	err := ctx.Bind(&req)
	if err != nil {
		return err
	}
	res, err := controller.ReportById(ctx, h.DB, *req)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, res)
}
