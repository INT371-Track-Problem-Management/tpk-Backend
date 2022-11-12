package controller

import (
	"net/http"
	"strconv"
	"tpk-backend/app/models/request"

	"github.com/labstack/echo/v4"
)

func (c controllerTPK) CreateReport(ctx echo.Context) error {
	req := new(request.ReportInsert)
	err := ctx.Bind(&req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	reportId, err := c.service.CreateReport(*req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, reportId)
}

func (c controllerTPK) ListReport(ctx echo.Context) error {
	var fillter = new(request.FillterReport)
	if roomId := ctx.QueryParam("roomId"); roomId != "" {
		fillter.RoomId = roomId
	}
	if customerId := ctx.QueryParam("customerId"); customerId != "" {
		fillter.CustomerId = customerId
	}
	reports, err := c.service.ListReport(fillter)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, reports)
}

func (c controllerTPK) ReportDetailById(ctx echo.Context) error {
	param := ctx.Param("reportId")
	reportId, err := strconv.Atoi(param)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	report, err := c.service.ReportDetailById(reportId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, report)
}

func (c controllerTPK) ReportEnagegeFixDateDetail(ctx echo.Context) error {
	reportId := ctx.Param("reportId")
	report, err := c.service.ReportEnagegeFixDateDetail(reportId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, report)
}

func (c controllerTPK) ChangeCategory(ctx echo.Context) error {
	req := new(request.ReportChangeCategory)
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	if err := c.service.ChangeCategory(*req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	response := map[string]string{
		"message": "success",
	}
	return ctx.JSON(http.StatusOK, response)
}

func (c controllerTPK) ChangeStatusReport(ctx echo.Context) error {
	req := new(request.ReportChangeStatus)
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	if err := c.service.ChangeStatusReport(*req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	response := map[string]string{
		"message": "success",
	}
	return ctx.JSON(http.StatusOK, response)
}

func (c controllerTPK) EndJobReport(ctx echo.Context) error {

	req := new(request.EndJobReport)
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	if err := c.service.EndJobReport(*req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	res := map[string]string{
		"message": "success",
	}

	return ctx.JSON(http.StatusOK, res)

}
