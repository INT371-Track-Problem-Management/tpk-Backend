package server

import (
	"net/http"
	"strconv"
	"tpk-backend/app/model/request"
	"tpk-backend/app/service/controller"

	"github.com/labstack/echo/v4"
)

// Owner service

func (h *FuncHandler) Rooms(ctx echo.Context) error {
	res, err := controller.Rooms(ctx, h.DB)
	if err != nil {

		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) RoomsStatus(ctx echo.Context) error {
	req := new(request.RoomsStatus)
	err := ctx.Bind(&req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	res, err := controller.RoomsStatus(ctx, h.DB, *req)
	if err != nil {

		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) Customer(ctx echo.Context) error {
	res, err := controller.Customer(ctx, h.DB)
	if err != nil {

		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) Dorm(ctx echo.Context) error {
	req := new(request.Dorm)
	err := ctx.Bind(&req)
	if err != nil {

		return ctx.JSON(http.StatusInternalServerError, err)
	}
	res, err := controller.Dorm(ctx, h.DB, *req)
	if err != nil {

		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) Report(ctx echo.Context) error {
	res, err := controller.Report(ctx, h.DB)
	if err != nil {

		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) ReportById(ctx echo.Context) error {
	req := new(request.Report)
	err := ctx.Bind(&req)
	if err != nil {

		return ctx.JSON(http.StatusBadRequest, err)
	}
	res, err := controller.ReportById(ctx, h.DB, *req)
	if err != nil {

		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) DormInsert(ctx echo.Context) error {
	req := new(request.DormInsert)
	err := ctx.Bind(&req)
	if err != nil {

		return ctx.JSON(http.StatusInternalServerError, err)
	}
	res, err := controller.DormInsert(ctx, h.DB, *req)
	if err != nil {

		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) RoomsInsert(ctx echo.Context) error {
	req := new(request.RoomInsert)
	err := ctx.Bind(&req)
	if err != nil {

		return ctx.JSON(http.StatusInternalServerError, err)
	}
	res, err := controller.RoomInsert(ctx, h.DB, *req)
	if err != nil {

		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) DormDelete(ctx echo.Context) error {
	req := new(request.DormDelete)
	err := ctx.Bind(&req)
	if err != nil {

		return ctx.JSON(http.StatusInternalServerError, err)
	}
	res, err := controller.DormDelete(ctx, h.DB, *req)
	if err != nil {

		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) ReportChangeStatus(ctx echo.Context) error {
	req := new(request.ReportChangeStatus)
	err := ctx.Bind(&req)
	if err != nil {

		return ctx.JSON(http.StatusInternalServerError, err)
	}
	res, err := controller.ReportChangeStatus(ctx, h.DB, *req)
	if err != nil {

		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) GetReportEngageAll(ctx echo.Context) error {
	param := ctx.QueryParam("dormId")
	id, _ := strconv.ParseInt(param, 10, 64)
	res, err := controller.GetReportEngageAll(ctx, h.DB, id)
	if err != nil {

		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) GetReportEngageById(ctx echo.Context) error {
	req := new(request.ReportEngageById)
	err := ctx.Bind(&req)
	if err != nil {

		return ctx.JSON(http.StatusInternalServerError, err)
	}
	res, err := controller.GetReportEngageById(ctx, h.DB, *req)
	if err != nil {

		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) InsertReportEngage(ctx echo.Context) error {
	req := new(request.ReportEngage)
	err := ctx.Bind(&req)
	if err != nil {

		return ctx.JSON(http.StatusInternalServerError, err)
	}
	res, err := controller.InsertReportEngage(ctx, h.DB, *req)
	if err != nil {

		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) DeleteReportById(ctx echo.Context) error {
	var err error
	req := new(request.Report)
	err = ctx.Bind(&req)
	if err != nil {

		return ctx.JSON(http.StatusBadRequest, err)
	}
	err = controller.DeleteReportById(ctx, h.DB, *req)
	if err != nil {

		return ctx.JSON(http.StatusInternalServerError, err)
	}
	res := map[string]string{
		"message": "success",
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) RoomByDormId(ctx echo.Context) error {
	dormId := ctx.QueryParam("dormId")
	res, err := controller.RoomByDormId(ctx, h.DB, dormId)
	if err != nil {

		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) GetCustomerById(ctx echo.Context) error {
	param := ctx.QueryParam("customerId")
	cusId, _ := strconv.ParseInt(param, 10, 32)
	res, err := controller.GetCustomerById(ctx, h.DB, cusId)
	if err != nil {

		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) EmployeeById(ctx echo.Context) error {
	param := ctx.QueryParam("employeeId")
	empId, _ := strconv.ParseInt(param, 10, 32)
	res, err := controller.EmployeeById(ctx, h.DB, empId)
	if err != nil {

		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) RoomAddCustomer(ctx echo.Context) error {
	var err error
	req := new(request.RoomAddCustomer)
	err = ctx.Bind(&req)
	if err != nil {

		return ctx.JSON(http.StatusInternalServerError, err)
	}
	err = controller.RoomAddCustomer(ctx, h.DB, *req)
	if err != nil {

		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, "sucess")
}

func (h *FuncHandler) RoomRemoveCustomer(ctx echo.Context) error {
	var err error
	param := ctx.QueryParam("id")
	id, _ := strconv.ParseInt(param, 10, 64)
	err = controller.RoomRemoveCustomer(ctx, h.DB, id)
	if err != nil {

		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, "sucess")
}

func (h *FuncHandler) GetAllRoomWithCustomer(ctx echo.Context) error {
	param := ctx.QueryParam("dormId")
	id, _ := strconv.ParseInt(param, 10, 64)
	res, err := controller.GetAllRoomWithCustomer(ctx, h.DB, id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) AddMaintainer(ctx echo.Context) error {
	req := new(request.Maintainer)
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	res, err := controller.AddMaintainer(ctx, h.DB, *req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	resdata := map[string]interface{}{
		"maintainerId": res,
		"message":      "success",
	}
	return ctx.JSON(http.StatusOK, resdata)
}

func (h *FuncHandler) CreateAssignFixReport(ctx echo.Context) error {
	req := new(request.AssignReport)
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	err := controller.CreateAssignFixReport(ctx, h.DB, *req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	res := map[string]string{
		"message": "success",
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) GetHistoryByEmployeeId(ctx echo.Context) error {
	param := ctx.QueryParam("employeeId")
	empId, _ := strconv.ParseInt(param, 10, 32)
	res, err := controller.GetHistoryByEmployeeId(ctx, h.DB, empId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) AddEmployeeInDorm(ctx echo.Context) error {
	req := new(request.AddEmpInDorm)
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	err := controller.AddEmployeeInDorm(ctx, h.DB, *req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	res := map[string]string{
		"message": "success",
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) GetReportEngageByReportId(ctx echo.Context) error {
	param := ctx.Param("reportId")
	reportId, _ := strconv.ParseInt(param, 10, 32)
	res, err := controller.GetReportEngageByReportId(ctx, h.DB, reportId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, res)
}
