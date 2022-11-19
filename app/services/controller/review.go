package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (c controllerTPK) FetchReviewByReportId(ctx echo.Context) error {
	reportId, _ := strconv.Atoi(ctx.Param("reportId"))
	report, err := c.service.FetchReviewByReportId(reportId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, report)
}
