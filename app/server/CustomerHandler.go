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
	id := ctx.QueryParam("cusid")
	err := authentication.ActivateCustomerCtr(ctx, h.DB, id, "A")
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	redir := URI + "login"
	fmt.Println("----test----")
	fmt.Println(redir)
	return ctx.Redirect(http.StatusMovedPermanently, redir)
}

func (h *FuncHandler) GetReportByCreatedBy(ctx echo.Context) error {
	req := new(request.ReportByCreatedBy)
	err := ctx.Bind(&req)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, err)
	}
	res, err := controller.GetReportByCreatedBy(ctx, h.DB, *req)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) GetCustomerProgfile(ctx echo.Context) error {
	email := ctx.QueryParam("email")
	res, err := controller.CustomerViewProfile(ctx, h.DB, email)
	if err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) CustomerEditProfile(ctx echo.Context) error {
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
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, &req)
}
func (h *FuncHandler) GetCustomerReportApplication(ctx echo.Context) error {
	req := new(request.Report)
	err := ctx.Bind(&req)
	res, err := controller.ReportById(ctx, h.DB, *req)
	if err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, res)

}

func (h *FuncHandler) FetchReportEngageJoinReport(ctx echo.Context) error {

	id := ctx.QueryParam("reportId")
	reportId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadGateway, err)
	}
	res, err := controller.ReportEngageJoinReport(ctx, h.DB, reportId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) SelectedPlanFixDate(ctx echo.Context) error {

	req := new(request.SelectedPlanFixDate)
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	if err := controller.SelectedDatePlanFix(ctx, h.DB, *req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	res := map[string]string{
		"massage": "success",
	}

	return ctx.JSON(http.StatusOK, res)

}

func (h *FuncHandler) EndJobReport(ctx echo.Context) error {

	req := new(request.EndJobReport)
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	if err := controller.EndJobReport(ctx, h.DB, *req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	res := map[string]string{
		"massage": "success",
	}

	return ctx.JSON(http.StatusOK, res)

}

func (h *FuncHandler) GetHistoryByCustomerId(ctx echo.Context) error {
	param := ctx.QueryParam("customerId")
	cusId, _ := strconv.ParseInt(param, 10, 32)
	res, err := controller.GetHistoryByCustomerId(ctx, h.DB, cusId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, res)
}
