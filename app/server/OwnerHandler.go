package server

import (
	"fmt"
	"net/http"
	"strconv"
	"tpk-backend/app/authentication"
	"tpk-backend/app/model/request"
	"tpk-backend/app/service/controller"

	"github.com/labstack/echo/v4"
)

// Owner service

func (h *FuncHandler) Rooms(ctx echo.Context) error {
	check, status := authentication.ValidateOwnerService(ctx)
	if status == false {
		return ctx.String(http.StatusUnauthorized, check.Token)
	}
	res, err := controller.Rooms(ctx, h.DB)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, "")
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) RoomsStatus(ctx echo.Context) error {
	check, status := authentication.ValidateOwnerService(ctx)
	if status == false {
		return ctx.String(http.StatusUnauthorized, check.Token)
	}
	req := new(request.RoomsStatus)
	err := ctx.Bind(&req)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, "")
	}
	res, err := controller.RoomsStatus(ctx, h.DB, *req)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, "")
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) Customer(ctx echo.Context) error {
	check, status := authentication.ValidateOwnerService(ctx)
	if status == false {
		return ctx.String(http.StatusUnauthorized, check.Token)
	}
	res, err := controller.Customer(ctx, h.DB)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, "")
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) Dorm(ctx echo.Context) error {
	check, status := authentication.ValidateOwnerService(ctx)
	if status == false {
		return ctx.String(http.StatusUnauthorized, check.Token)
	}
	req := new(request.Dorm)
	err := ctx.Bind(&req)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, "")
	}
	res, err := controller.Dorm(ctx, h.DB, *req)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, "")
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) Report(ctx echo.Context) error {
	check, status := authentication.ValidateOwnerService(ctx)
	if status == false {
		return ctx.String(http.StatusUnauthorized, check.Token)
	}
	res, err := controller.Report(ctx, h.DB)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, "")
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) ReportById(ctx echo.Context) error {
	check, status := authentication.ValidateOwnerService(ctx)
	if status == false {
		return ctx.String(http.StatusUnauthorized, check.Token)
	}
	req := new(request.Report)
	err := ctx.Bind(&req)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, "")
	}
	res, err := controller.ReportById(ctx, h.DB, *req)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, "")
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) DormInsert(ctx echo.Context) error {
	check, status := authentication.ValidateOwnerService(ctx)
	if status == false {
		return ctx.String(http.StatusUnauthorized, check.Token)
	}
	req := new(request.DormInsert)
	err := ctx.Bind(&req)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, "")
	}
	res, err := controller.DormInsert(ctx, h.DB, *req)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, "")
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) RoomsInsert(ctx echo.Context) error {
	check, status := authentication.ValidateOwnerService(ctx)
	if status == false {
		return ctx.String(http.StatusUnauthorized, check.Token)
	}
	req := new(request.RoomInsert)
	err := ctx.Bind(&req)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, "")
	}
	res, err := controller.RoomInsert(ctx, h.DB, *req)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, "")
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) DormDelete(ctx echo.Context) error {
	check, status := authentication.ValidateOwnerService(ctx)
	if status == false {
		return ctx.String(http.StatusUnauthorized, check.Token)
	}
	req := new(request.DormDelete)
	err := ctx.Bind(&req)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, "")
	}
	res, err := controller.DormDelete(ctx, h.DB, *req)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, "")
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) ReportChangeStatus(ctx echo.Context) error {
	check, status := authentication.ValidateOwnerService(ctx)
	if status == false {
		return ctx.String(http.StatusUnauthorized, check.Token)
	}
	req := new(request.ReportChangeStatus)
	err := ctx.Bind(&req)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, "")
	}
	res, err := controller.ReportChangeStatus(ctx, h.DB, *req)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, "")
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) GetReportEngageAll(ctx echo.Context) error {
	check, status := authentication.ValidateOwnerService(ctx)
	if status == false {
		return ctx.String(http.StatusUnauthorized, check.Token)
	}
	res, err := controller.GetReportEngageAll(ctx, h.DB)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, "")
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) GetReportEngageById(ctx echo.Context) error {
	check, status := authentication.ValidateOwnerService(ctx)
	if status == false {
		return ctx.String(http.StatusUnauthorized, check.Token)
	}
	req := new(request.ReportEngageById)
	err := ctx.Bind(&req)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, "")
	}
	res, err := controller.GetReportEngageById(ctx, h.DB, *req)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, "")
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) InsertReportEngage(ctx echo.Context) error {
	check, status := authentication.ValidateOwnerService(ctx)
	if status == false {
		return ctx.String(http.StatusUnauthorized, check.Token)
	}
	req := new(request.ReportEngage)
	err := ctx.Bind(&req)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, "")
	}
	res, err := controller.InsertReportEngage(ctx, h.DB, *req)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, "")
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) DeleteReportById(ctx echo.Context) error {
	check, status := authentication.ValidateOwnerService(ctx)
	if status == false {
		return ctx.String(http.StatusUnauthorized, check.Token)
	}
	var err error
	req := new(request.Report)
	err = ctx.Bind(&req)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, "")
	}
	err = controller.DeleteReportById(ctx, h.DB, *req)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, "")
	}
	return ctx.JSON(http.StatusNoContent, "")
}

func (h *FuncHandler) RoomByDormId(ctx echo.Context) error {
	check, status := authentication.ValidateOwnerService(ctx)
	if status == false {
		return ctx.String(http.StatusUnauthorized, check.Token)
	}
	dormId := ctx.QueryParam("dormId")
	res, err := controller.RoomByDormId(ctx, h.DB, dormId)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, "")
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) ReportByDormId(ctx echo.Context) error {
	check, status := authentication.ValidateOwnerService(ctx)
	if status == false {
		return ctx.String(http.StatusUnauthorized, check.Token)
	}
	dormId := ctx.QueryParam("dormId")
	res, err := controller.ReportByDormId(ctx, h.DB, dormId)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, "")
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) GetCustomerById(ctx echo.Context) error {
	check, status := authentication.ValidateOwnerService(ctx)
	if status == false {
		return ctx.String(http.StatusUnauthorized, check.Token)
	}
	param := ctx.QueryParam("customerId")
	cusId, _ := strconv.ParseInt(param, 10, 32)
	res, err := controller.GetCustomerById(ctx, h.DB, cusId)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, "")
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) EmployeeById(ctx echo.Context) error {
	check, status := authentication.ValidateOwnerService(ctx)
	if status == false {
		return ctx.String(http.StatusUnauthorized, check.Token)
	}
	param := ctx.QueryParam("employeeId")
	empId, _ := strconv.ParseInt(param, 10, 32)
	res, err := controller.EmployeeById(ctx, h.DB, empId)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, "")
	}
	return ctx.JSON(http.StatusOK, res)
}
