package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"tpk-backend/app/config"
	"tpk-backend/app/jwt"
	"tpk-backend/app/services/controller"
	"tpk-backend/app/services/repository"
	"tpk-backend/app/services/service"
	"tpk-backend/app/services/validator"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func StartServer() {
	log.Println(config.LoadTest())
	key := os.Getenv("KEY")
	port := SetEnv(key)
	fmt.Println("PROJECT RUN ON PORT: " + port)
	e := echo.New()
	db, _ := InitializeDatabase()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*", "localhost"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	repo := repository.NewRepository(db)
	serviceTPK := service.NewService(repo, db)
	controller := controller.NewController(serviceTPK)
	validator := validator.Newvalidator(db)

	api := e.Group("/api/")
	api.GET("checkHealthy", controller.CheckHealthy)
	api.POST("login", controller.Login)
	api.POST("registerCustomer", controller.RegisterCustomers) // Register customer
	api.POST("registerOwner", controller.RegisterOwner)        // Register owner

	service := api.Group("service/")
	service.Use(middleware.JWTWithConfig(jwt.ValidateTokenJWTConfig()))

	cus := api.Group("customer/")
	cus.Use(middleware.JWTWithConfig(jwt.ValidateTokenJWTConfig()))
	cus.POST("report", controller.CreateReport, validator.CustomerValidation)                                 // Insert report
	cus.GET("reports", controller.ListReport, validator.CustomerValidation)                                   // Search all report
	cus.GET("report/detail/:reportId", controller.ReportDetailById, validator.CustomerValidation)             // Search report detail by id
	cus.GET("reportStatus/detail/:reportId", controller.ReportStatusByReportId, validator.CustomerValidation) // Search reportStatus detail by reportId

	emp := api.Group("employee/")
	emp.Use(middleware.JWTWithConfig(jwt.ValidateTokenJWTConfig()))
	emp.GET("reports", controller.ListReport, validator.EmployeeValidation)                                   // Search all report
	emp.GET("report/detail/:reportId", controller.ReportDetailById, validator.EmployeeValidation)             // Search report detail by id
	emp.GET("reportStatus/detail/:reportId", controller.ReportStatusByReportId, validator.EmployeeValidation) // Search reportStatus detail by reportId

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
	if key == "" {
		port = "3050"
		URI = "http://localhost:3050/"
		URIRedi = "https://dev.rungmod.com/"
		return port
	} else {
		port = "3050"
		URI = "http://localhost:3050/"
		log.Println("Invalid port ENV")
	}
	return ""
}
