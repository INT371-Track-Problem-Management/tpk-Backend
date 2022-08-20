package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"tpk-backend/app/authentication"
	"tpk-backend/app/pkg/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func StartServer() {

	log.Println(config.LoadTest())
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
	api.GET("testEmail", h.TestGmail)
	api.POST("registerCustomer", h.RegisterCustomer) // Register customer
	api.POST("registerOwner", h.RegisterOwner)       // Register owner

	// Both but need TOKEN
	service := api.Group("service/")
	service.Use(middleware.JWTWithConfig(authentication.ValidateTokenJWTConfig()))
	service.POST("changeEmail", h.ChangeEmail) // Change email customer or employee
	service.GET("decodeRole", h.GetRoleJWT)    // Decode TOken to get Role

	// Customer Service
	cus := api.Group("customer/")
	cus.Use(middleware.JWTWithConfig(authentication.ValidateTokenJWTConfig()))
	cus.GET("checkHealthy", h.CheckHealthyJWT)                            //Check Heatkhy with Token
	cus.GET("activateCus", h.ActivateCustomer)                            // Activate user change status 'I' => 'A'
	cus.GET("reportByCreatedBy", h.GetReportByCreatedBy)                  // Search report by createdBy
	cus.POST("report", h.ReportInsert)                                    // Insert report
	cus.GET("viewCustomerProfile/*", h.GetCustomerProgfile)               // View profile customer by email
	cus.PUT("editProfile/*", h.CustomerEditProfile)                       // Edit customer profile
	cus.GET("getReportEngageWithReport/*", h.FetchReportEngageJoinReport) // Seach reportEngage join with reports whare by customerId
	cus.GET("selectedPlanFixDate", h.SelectedPlanFixDate)                 // customer selecting plan fix date

	// Owner Service
	emp := api.Group("employee/")
	emp.Use(middleware.JWTWithConfig(authentication.ValidateTokenJWTConfig()))
	emp.POST("reportEngageById", h.GetReportEngageById)           // Search Report by engageId
	emp.POST("CreateReportEngage", h.InsertReportEngage)          // Insert Report Engage
	emp.PUT("statusReport", h.ReportChangeStatus)                 // Update status Report
	emp.DELETE("deleteReportById", h.DeleteReportById)            // Delete report by Id
	emp.GET("rooms", h.Rooms)                                     // Search all room
	emp.GET("customer", h.Customer)                               // Search all customer
	emp.PUT("rooms", h.RoomsStatus)                               // Change room status
	emp.GET("dorm", h.Dorm)                                       // Search all dorm
	emp.GET("report", h.Report)                                   // Search all report
	emp.POST("reportById", h.ReportById)                          // Search report by id
	emp.POST("dorm", h.DormInsert)                                // Insert Dorm
	emp.POST("rooms", h.RoomsInsert)                              // Insert Room
	emp.DELETE("dorm", h.DormDelete)                              // Delete dorm
	emp.GET("reportEngageAll", h.GetReportEngageAll)              // Search all report engage
	emp.GET("reportByDormId/*", h.ReportByDormId)                 // Search report by dormId
	emp.GET("roomByDormId/*", h.RoomByDormId)                     // Search room by dormId
	emp.GET("customerById/*", h.GetCustomerById)                  // Search customer by Id
	emp.GET("employeeById/*", h.EmployeeById)                     // Search rmployee by Id
	emp.POST("roomAddCustomer", h.RoomAddCustomer)                // Add customer into room and room status 'I'=> 'A'
	emp.GET("GetAllRoomWithCustomer/*", h.GetAllRoomWithCustomer) // Search all customer in their dormId

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

type FuncHandler struct {
	DB *gorm.DB
}
