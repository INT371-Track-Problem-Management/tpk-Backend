package server

import (
	"fmt"
	"net/http"
	"os"
	"tpk-backend/app/authentication"
	"tpk-backend/app/model/request"
	"tpk-backend/app/pkg"
	"tpk-backend/app/pkg/config"
	"tpk-backend/app/service/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	api := e.Group("/api/")
	api.POST("login", h.Login)
	api.GET("test", h.Test)
	api.GET("checkHealthy", h.CheckHealthy)
	api.GET("rooms", h.Rooms)
	api.GET("customer", h.Customer)
	api.PUT("rooms", h.RoomsStatus)
	api.GET("dorm", h.Dorm)
	api.GET("report", h.Report)
	api.GET("reportById", h.ReportById)
	api.POST("dorm", h.DormInsert)
	api.POST("rooms", h.RoomsInsert)
	api.DELETE("dorm", h.DormDelete)
	api.POST("report", h.ReportInsert)
	api.PUT("statusReport", h.ReportChangeStatus)
	api.GET("testEmail", h.TestGmail)
	api.POST("registerCustomer", h.RegisterCustomer)
	api.GET("reportEngageAll", h.GetReportEngageAll)
	api.GET("reportEngageById", h.GetReportEngageById)
	api.POST("CreateReportEngage", h.InsertReportEngage)
	api.GET("activateCus", h.ActivateCustomer)

	e.Logger.Fatal(e.Start(":" + port))
}

var URI string
var URIRedi string

func SetEnv(key string) string {
	var port string
	if key == "PRD" {
		port = "5000"
		URI = "https://www.rungmod.com/"
		return port
	}
	if key == "DEV" {
		port = "3000"
		URI = "https://dev.rungmod.com/"
		return port
	}
	if key == "local" {
		port = "3050"
		URI = "http://localhost:3050/"
		URIRedi = "https://dev.rungmod.com/"
		return port
	} else {
		fmt.Printf("Invalid ENV")
	}
	return ""
}

type FuncHandler struct {
	DB *gorm.DB
}

func (h *FuncHandler) Login(ctx echo.Context) error {
	var token *string
	var err error
	user := new(request.User)
	err = ctx.Bind(&user)
	if err != nil {
		return err
	}
	token, err = authentication.Login(ctx, h.DB, *user)
	if err != nil {
		fmt.Println("Unatutherize")
		return ctx.JSON(http.StatusUnauthorized, "Unatutherize")
	}
	return ctx.JSON(http.StatusOK, token)
}

func (h *FuncHandler) Initialize() {
	db := config.LoadDB()
	dns := fmt.Sprintf(`%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local`, db.Username, db.Password, db.Host, db.Port, db.Database)
	conn, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println(err)
	}
	h.DB = conn
}
func (h *FuncHandler) TestGmail(ctx echo.Context) error {
	testmail := "artid.vijitpanmai@mail.kmutt.ac.th"
	pkg.SSLemail(&testmail, "Hello-World", "Hi")
	return ctx.JSON(http.StatusOK, "send mail success")
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
	fmt.Println("test")
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

func (h *FuncHandler) DormInsert(ctx echo.Context) error {
	req := new(request.DormInsert)
	err := ctx.Bind(&req)
	if err != nil {
		return err
	}
	res, err := controller.DormInsert(ctx, h.DB, *req)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) RoomsInsert(ctx echo.Context) error {
	req := new(request.RoomInsert)
	err := ctx.Bind(&req)
	if err != nil {
		return err
	}
	res, err := controller.RoomInsert(ctx, h.DB, *req)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) DormDelete(ctx echo.Context) error {
	req := new(request.DormDelete)
	err := ctx.Bind(&req)
	if err != nil {
		return err
	}
	res, err := controller.DormDelete(ctx, h.DB, *req)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) ReportInsert(ctx echo.Context) error {
	req := new(request.ReportInsert)
	err := ctx.Bind(&req)
	if err != nil {
		return err
	}
	res, err := controller.ReportInsert(ctx, h.DB, *req)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) ReportChangeStatus(ctx echo.Context) error {
	req := new(request.ReportChangeStatus)
	err := ctx.Bind(&req)
	if err != nil {
		return err
	}
	res, err := controller.ReportChangeStatus(ctx, h.DB, *req)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) RegisterCustomer(ctx echo.Context) error {
	var err error
	req := new(authentication.RegisterCustomer)
	err = ctx.Bind(&req)
	if err != nil {
		return err
	}
	customerId, err := authentication.RegisterCustomers(ctx, h.DB, *req, URI)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, customerId)
}

func (h *FuncHandler) GetReportEngageAll(ctx echo.Context) error {
	res, err := controller.GetReportEngageAll(ctx, h.DB)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) GetReportEngageById(ctx echo.Context) error {
	req := new(request.ReportEngageById)
	err := ctx.Bind(&req)
	if err != nil {
		return err
	}
	res, err := controller.GetReportEngageById(ctx, h.DB, *req)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) InsertReportEngage(ctx echo.Context) error {
	req := new(request.ReportEngage)
	err := ctx.Bind(&req)
	if err != nil {
		return err
	}
	res, err := controller.InsertReportEngage(ctx, h.DB, *req)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) ActivateCustomer(ctx echo.Context) error {
	id := ctx.QueryParam("cusid")
	err := authentication.ActivateCustomerCtr(ctx, h.DB, id, "A")
	if err != nil {
		return err
	}
	redir := URI + "login"
	fmt.Println("----test----")
	fmt.Println(redir)
	return ctx.Redirect(http.StatusMovedPermanently, redir)
}
