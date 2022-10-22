package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"tpk-backend/app/authentication"
	"tpk-backend/app/pkg/config"
	"tpk-backend/app/validator"

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
		AllowOrigins: []string{"*", "localhost"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	api := e.Group("/api/")
	api.POST("login", h.Login)
	api.POST("logout", h.Logout)
	api.GET("test", h.Test)
	api.GET("checkHealthy", h.CheckHealthy)
	api.POST("registerCustomer", h.RegisterCustomer) // Register customer
	api.POST("registerOwner", h.RegisterOwner)       // Register owner
	api.GET("yearConfig", h.YearConfig)              // Get all year between 1901 - 2022
	api.GET("testemail", h.TestSMTP2)

	// Both but need TOKEN
	service := api.Group("service/")
	service.Use(middleware.JWTWithConfig(authentication.ValidateTokenJWTConfig()))
	service.POST("changeEmail", h.ChangeEmail)                    // Change email customer or employee
	service.GET("decodeRole", h.GetRoleJWT)                       // Decode TOken to get Role
	service.GET("historyreportById/*", h.GetHistoryByHistoryId)   // Search history by Id
	service.POST("changePassword", h.ChangePassword)              // Change password
	service.GET("maintainerById/:maintainerId", h.MaintainerById) // Search maintainer by Id

	// Customer Service
	cus := api.Group("customer/")
	cus.Use(middleware.JWTWithConfig(authentication.ValidateTokenJWTConfig()))
	cus.GET("checkHealthy", h.CheckHealthyJWT, validator.CustomerValidation)                                   //Check Heatkhy with Token
	cus.GET("activateCus", h.ActivateCustomer, validator.CustomerValidation)                                   // Activate user change status 'I' => 'A'
	cus.GET("reportByCreatedBy/:customerId", h.GetReportByCreatedBy, validator.CustomerValidation)             // Search report by createdBy
	cus.POST("report", h.ReportInsert, validator.CustomerValidation)                                           // Insert report
	cus.GET("viewCustomerProfile/*", h.GetCustomerProgfile, validator.CustomerValidation)                      // View profile customer by email
	cus.PUT("editProfile/*", h.CustomerEditProfile, validator.CustomerValidation)                              // Edit customer profile
	cus.GET("getReportEngageWithReport/*", h.FetchReportEngageJoinReport, validator.CustomerValidation)        // Seach reportEngage join with reports whare by customerId
	cus.PUT("selectedPlanFixDate", h.SelectedPlanFixDate, validator.CustomerValidation)                        // customer selecting plan fix date
	cus.POST("endJobReview", h.EndJobReport, validator.CustomerValidation)                                     // end job report and review
	cus.GET("historyReport/list/*", h.GetHistoryByCustomerId, validator.CustomerValidation)                    // Search all history by customerId
	cus.POST("reportById", h.ReportById, validator.CustomerValidation)                                         // Search report by id
	cus.GET("getAllRoomByCustomerId/:customerId", h.GetRoomsByCustomerId, validator.CustomerValidation)        // search all rooms by customerId
	cus.GET("getAllReportByRoomId/:roomId", h.ReportByRoomId, validator.CustomerValidation)                    // search reports by roomId
	cus.GET("FetchProfile/:email", h.FetchProfile, validator.CustomerValidation)                               // Search maintainer by Id
	cus.GET("reportStatusApplication/:reportId", h.FetchReportStatusApplication, validator.CustomerValidation) // Find all status report by report Id
	cus.GET("listReports/:customerId", h.ReportListForCustomer, validator.CustomerValidation)                  // list report by customer id
	cus.GET("getAllReportByRoomId/:roomId", h.ReportByRoomId, validator.CustomerValidation)                    // search reports by roomId

	// Owner Service
	emp := api.Group("employee/")
	emp.Use(middleware.JWTWithConfig(authentication.ValidateTokenJWTConfig()))
	emp.POST("reportEngageById", h.GetReportEngageById, validator.EmployeeValidation)  // Search Report by engageId
	emp.POST("CreateReportEngage", h.InsertReportEngage, validator.EmployeeValidation) // Insert Report Engage
	emp.PUT("statusReport", h.ReportChangeStatus, validator.EmployeeValidation)        // Update status Report
	emp.DELETE("deleteReportById", h.DeleteReportById, validator.EmployeeValidation)   // Delete report by Id
	emp.GET("rooms", h.Rooms, validator.EmployeeValidation)                            // Search all room
	emp.GET("roomByRoomNum/*", h.RoomByRoomNum, validator.EmployeeValidation)          // Search room by roomNum
	emp.GET("roomByRoomId/*", h.RoomByRoomId, validator.EmployeeValidation)            // Search room by roomId
	emp.GET("rooms", h.Rooms, validator.EmployeeValidation)
	emp.GET("customer", h.Customer, validator.EmployeeValidation)                                              // Search all customer
	emp.PUT("rooms", h.RoomsStatus, validator.EmployeeValidation)                                              // Change room status
	emp.GET("buildingById/:buildingId", h.BuildingById, validator.EmployeeValidation)                          // Search building by Id
	emp.GET("allBuilding", h.AllBuilding, validator.EmployeeValidation)                                        // Search all building
	emp.GET("report", h.Report, validator.EmployeeValidation)                                                  // Search all report
	emp.POST("reportById", h.ReportById, validator.EmployeeValidation)                                         // Search report by id
	emp.POST("building", h.BuildingInsert, validator.EmployeeValidation)                                       // Insert Building
	emp.POST("rooms", h.RoomsInsert, validator.EmployeeValidation)                                             // Insert Room
	emp.DELETE("building", h.BuildingDelete, validator.EmployeeValidation)                                     // Delete building
	emp.GET("reportEngageAll/*", h.GetReportEngageAll, validator.EmployeeValidation)                           // Search all report engage
	emp.GET("roomByBuildingId/:buildingId", h.RoomByBuildingId, validator.EmployeeValidation)                  // Search room by buildingId
	emp.GET("customerById/*", h.GetCustomerById, validator.EmployeeValidation)                                 // Search customer by Id
	emp.GET("employeeById/*", h.EmployeeById, validator.EmployeeValidation)                                    // Search rmployee by Id
	emp.POST("roomAddCustomer", h.RoomAddCustomer, validator.EmployeeValidation)                               // Add customer into room and room status 'I'=> 'A'
	emp.GET("GetAllRoomWithCustomer/*", h.GetAllRoomWithCustomer, validator.EmployeeValidation)                // Search all customer in their buildingId
	emp.POST("maintainer", h.AddMaintainer, validator.EmployeeValidation)                                      // Created maintainer and return Id
	emp.POST("assignFixReport", h.CreateAssignFixReport, validator.EmployeeValidation)                         // add maintainer to fix report
	emp.GET("historyReport/list/*", h.GetHistoryByEmployeeId, validator.EmployeeValidation)                    // Search all history by employeeId
	emp.POST("addEmployeeInBuilding", h.AddEmployeeInBuilding, validator.EmployeeValidation)                   // Add employee in building and change position to staff
	emp.GET("reportEngageByReportId/:reportId", h.GetReportEngageByReportId, validator.EmployeeValidation)     // Search reportEngage by reportId
	emp.POST("dashboard", h.FetcStatDashBoard, validator.EmployeeValidation)                                   // Get stat for dashboard
	emp.GET("getAllReportByRoomId/:roomId", h.ReportByRoomId, validator.EmployeeValidation)                    // search reports by roomId
	emp.GET("maintainerList", h.Maintainerlist, validator.EmployeeValidation)                                  // Search all maintainer
	emp.GET("FetchProfile/:email", h.FetchProfile, validator.EmployeeValidation)                               // Search maintainer by Id
	emp.GET("reportStatusApplication/:reportId", h.FetchReportStatusApplication, validator.CustomerValidation) // Find all status report by report Id

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
