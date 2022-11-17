package controller

import (
	"net/http"
	"tpk-backend/app/config"
	"tpk-backend/app/pkg"

	"github.com/labstack/echo/v4"
)

func (c controllerTPK) TestUploadFile(ctx echo.Context) error {
	file, err := ctx.FormFile("image")
	if err != nil {
		return err
	}

	image, err := pkg.UploadReportFile(file)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	response := map[string]interface{}{
		"image": *image,
	}

	return ctx.JSON(http.StatusOK, response)
}

func (c controllerTPK) DownloadReportImage(ctx echo.Context) error {
	id := ctx.Param("image_id")
	image, err := c.service.ReportMediaById(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	pathImage := config.LoadPathMedia()
	path := pathImage.Path + "report/" + image.FileName
	return ctx.File(path)
}
