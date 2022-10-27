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

func (h *FuncHandler) BuildingById(ctx echo.Context) error {
	buildingId := ctx.Param("buildingId")
	res, err := controller.BuildingById(ctx, h.DB, buildingId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) AllBuilding(ctx echo.Context) error {
	data, err := controller.AllBuilding(ctx, h.DB)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	res := map[string]interface{}{
		"AllBuilding": data,
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

func (h *FuncHandler) BuildingInsert(ctx echo.Context) error {
	req := new(request.BuildingInsert)
	err := ctx.Bind(&req)
	if err != nil {

		return ctx.JSON(http.StatusInternalServerError, err)
	}
	id, err := controller.BuildingInsert(ctx, h.DB, *req)
	if err != nil {

		return ctx.JSON(http.StatusInternalServerError, err)
	}
	res := map[string]interface{}{
		"message":    "Insert success",
		"BuildingId": id,
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) RoomsInsert(ctx echo.Context) error {
	req := new(request.RoomInsert)
	err := ctx.Bind(&req)
	if err != nil {

		return ctx.JSON(http.StatusInternalServerError, err)
	}
	err = controller.RoomInsert(ctx, h.DB, *req)
	if err != nil {
		errmsg := map[string]interface{}{
			"message":           "Can not create rooms",
			"error description": err,
		}
		return ctx.JSON(http.StatusInternalServerError, errmsg)
	}
	res := map[string]string{
		"message": "success",
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) BuildingDelete(ctx echo.Context) error {
	req := new(request.BuildingDelete)
	err := ctx.Bind(&req)
	if err != nil {

		return ctx.JSON(http.StatusInternalServerError, err)
	}
	res, err := controller.BuildingDelete(ctx, h.DB, *req)
	if err != nil {

		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) GetReportEngageAll(ctx echo.Context) error {
	param := ctx.QueryParam("buildingId")
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

func (h *FuncHandler) RoomByBuildingId(ctx echo.Context) error {
	buildingId := ctx.Param("buildingId")
	res, err := controller.RoomByBuildingId(ctx, h.DB, buildingId)
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
	return ctx.JSON(http.StatusOK, "success")
}

func (h *FuncHandler) RoomRemoveCustomer(ctx echo.Context) error {
	var err error
	param := ctx.QueryParam("id")
	id, _ := strconv.ParseInt(param, 10, 64)
	err = controller.RoomRemoveCustomer(ctx, h.DB, id)
	if err != nil {

		return ctx.JSON(http.StatusInternalServerError, err)
	}
	res := map[string]string{
		"message": "success",
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) GetAllRoomWithCustomer(ctx echo.Context) error {
	param := ctx.QueryParam("buildingId")
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

func (h *FuncHandler) AddEmployeeInBuilding(ctx echo.Context) error {
	req := new(request.AddEmpInBuilding)
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	err := controller.AddEmployeeInBuilding(ctx, h.DB, *req)
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

func (h *FuncHandler) RoomByRoomId(ctx echo.Context) error {
	roomId := ctx.QueryParam("roomId")
	res, err := controller.RoomByRoomId(ctx, h.DB, roomId)
	if err != nil {

		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) RoomByRoomNum(ctx echo.Context) error {
	roomNum := ctx.QueryParam("roomNum")
	res, err := controller.RoomByRoomNum(ctx, h.DB, roomNum)
	if err != nil {

		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) FetcStatDashBoard(ctx echo.Context) error {
	req := new(request.Stat)
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	res, err := controller.FetcStatDashBoard(ctx, h.DB, *req)
	if err != nil {

		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) Maintainerlist(ctx echo.Context) error {
	res, err := controller.Maintainerlist(ctx, h.DB)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) ListEmployee(ctx echo.Context) error {
	employees, err := controller.ListEmployee(ctx, h.DB)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	response := map[string]interface{}{
		"Employees": employees,
	}
	return ctx.JSON(http.StatusOK, response)
}
