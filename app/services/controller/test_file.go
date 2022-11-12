package controller

import (
	"net/http"
	"tpk-backend/app/pkg"

	"github.com/labstack/echo/v4"
)

func (c controllerTPK) TestUploadFile(ctx echo.Context) error {
	file, err := ctx.FormFile("image")
	if err != nil {
		return err
	}

	filename, err := pkg.UploadFile(file)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	response := map[string]string{
		"image_name": *filename,
	}

	return ctx.JSON(http.StatusOK, response)
}

func (c controllerTPK) TestDownloadFile(ctx echo.Context) error {
	name := ctx.Param("image_name")
	path := "../../images/" + name
	return ctx.File(path)
}
