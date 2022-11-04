package controller

import (
	"fmt"
	"net/http"
	"tpk-backend/app/models/request"

	"github.com/labstack/echo/v4"
)

func (c controllerTPK) RegisterCustomers(ctx echo.Context) error {
	var err error
	req := new(request.RegisterCustomer)
	err = ctx.Bind(&req)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	cusId, err := c.service.RegisterCustomersService(*req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, cusId)
}

func (c controllerTPK) RegisterOwner(ctx echo.Context) error {
	var err error
	req := new(request.RegisterOwner)
	err = ctx.Bind(&req)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	empId, err := c.service.RegisterOwnerService(*req)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, empId)
}
