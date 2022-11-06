package controller

import (
	"net/http"
	"tpk-backend/app/models/request"

	"github.com/labstack/echo/v4"
)

func (c controllerTPK) CreateReport(ctx echo.Context) error {
	req := new(request.ReportInsert)
	err := ctx.Bind(&req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	reportId, err := c.service.CreateReport(*req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	response := map[string]interface{}{
		"reportId": reportId,
	}
	return ctx.JSON(http.StatusBadRequest, response)
}
