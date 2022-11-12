package controller

import (
	"net/http"
	"tpk-backend/app/constants"
	"tpk-backend/app/pkg"

	"github.com/labstack/echo/v4"
)

func (c controllerTPK) TestUploadFile(ctx echo.Context) error {
	file, err := ctx.FormFile("image")
	if err != nil {
		return err
	}

	image, err := pkg.UploadFile(file, constants.IMAGE_DES_REPORT)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	response := map[string]interface{}{
		"image": *image,
	}

	return ctx.JSON(http.StatusOK, response)
}

func (c controllerTPK) TestDownloadFile(ctx echo.Context) error {
	name := ctx.Param("image_name")
	path := "../../images/report/" + name
	return ctx.File(path)
}
