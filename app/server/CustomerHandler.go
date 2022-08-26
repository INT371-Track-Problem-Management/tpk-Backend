package server

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"tpk-backend/app/authentication"
	"tpk-backend/app/model/request"
	"tpk-backend/app/service/controller"

	"github.com/labstack/echo/v4"
)

// Customer service

func (h *FuncHandler) ReportInsert(ctx echo.Context) error {
	check, status := authentication.ValidateCustomerService(ctx)
	if status == false {
		return ctx.String(http.StatusUnauthorized, check.Token)
	}
	req := new(request.ReportInsert)
	err := ctx.Bind(&req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	log.Println(req)
	res, err := controller.ReportInsert(ctx, h.DB, *req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) ActivateCustomer(ctx echo.Context) error {
	check, status := authentication.ValidateCustomerService(ctx)
	if status == false {
		return ctx.String(http.StatusUnauthorized, check.Token)
	}
	id := ctx.QueryParam("cusid")
	err := authentication.ActivateCustomerCtr(ctx, h.DB, id, "A")
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, "")
	}
	redir := URI + "login"
	fmt.Println("----test----")
	fmt.Println(redir)
	return ctx.Redirect(http.StatusMovedPermanently, redir)
}

func (h *FuncHandler) GetReportByCreatedBy(ctx echo.Context) error {
	check, status := authentication.ValidateCustomerService(ctx)
	if status == false {
		return ctx.String(http.StatusUnauthorized, check.Token)
	}
	req := new(request.ReportByCreatedBy)
	err := ctx.Bind(&req)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, "")
	}
	res, err := controller.GetReportByCreatedBy(ctx, h.DB, *req)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, "")
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) GetCustomerProgfile(ctx echo.Context) error {
	check, status := authentication.ValidateCustomerService(ctx)
	if status == false {
		return ctx.String(http.StatusUnauthorized, check.Token)
	}
	email := ctx.QueryParam("email")
	res, err := controller.CustomerViewProfile(ctx, h.DB, email)
	if err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) CustomerEditProfile(ctx echo.Context) error {
	check, status := authentication.ValidateCustomerService(ctx)
	if status == false {
		return ctx.String(http.StatusUnauthorized, check.Token)
	}
	req := new(request.CustomerEditProfile)
	err := ctx.Bind(&req)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, err)
	}
	email := ctx.QueryParam("email")
	err = controller.CustomerEditProfile(ctx, h.DB, *req, email)
	if err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, &req)
}
func (h *FuncHandler) GetCustomerReportApplication(ctx echo.Context) error {
	req := new(request.Report)
	err := ctx.Bind(&req)
	res, err := controller.ReportById(ctx, h.DB, *req)
	if err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusBadRequest, err)
	}
	check, status := authentication.ValidateCustomerService(ctx)
	if status == false {
		return ctx.String(http.StatusUnauthorized, check.Token)
	}

	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, res)

}

func (h *FuncHandler) FetchReportEngageJoinReport(ctx echo.Context) error {
	check, status := authentication.ValidateCustomerService(ctx)
	if status == false {
		return ctx.String(http.StatusUnauthorized, check.Token)
	}

	id := ctx.QueryParam("customer_id")
	customerId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadGateway, err)
	}
	res, err := controller.ReportEngageJoinReport(ctx, h.DB, customerId)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) SelectedPlanFixDate(ctx echo.Context) error {
	if check, status := authentication.ValidateCustomerService(ctx); status == false {
		return ctx.String(http.StatusUnauthorized, check.Token)
	}

	req := new(request.SelectedPlanFixDate)
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	if err := controller.SelectedDatePlanFix(ctx, h.DB, *req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	res := map[string]string{
		"massage": "success",
	}

	return ctx.JSON(http.StatusOK, res)

}
