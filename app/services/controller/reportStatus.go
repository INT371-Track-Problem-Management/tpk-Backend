package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c controllerTPK) ReportStatusByReportId(ctx echo.Context) error {
	reportId := ctx.Param("reportId")
	report, err := c.service.ReportStatusByReportId(reportId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, report)
}
