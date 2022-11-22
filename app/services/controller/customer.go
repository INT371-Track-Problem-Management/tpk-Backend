package controller

import (
	"net/http"
	"strconv"
	"tpk-backend/app/jwt"

	"github.com/labstack/echo/v4"
)

func (c controllerTPK) FetchCustomerByEmail(ctx echo.Context) error {
	user := jwt.DecodeJWT(ctx)

	customer, err := c.service.FetchCustomerByEmail(user.Email)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	respone := map[string]interface{}{
		"customer": customer,
	}
	return ctx.JSON(http.StatusOK, respone)
}

func (c controllerTPK) GetListCustomer(ctx echo.Context) error {
	customers, err := c.service.GetListCustomer()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	if len(customers) == 0 {
		return ctx.JSON(http.StatusNoContent, "No content")
	}

	response := map[string]interface{}{
		"customers": customers,
	}
	return ctx.JSON(http.StatusOK, response)
}

func (c controllerTPK) FetchCustomerById(ctx echo.Context) error {
	customerId, _ := strconv.Atoi(ctx.Param("customerId"))
	customer, err := c.service.FetchCustomerById(customerId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	response := map[string]interface{}{
		"customer": customer,
	}
	return ctx.JSON(http.StatusOK, response)
}

func (c controllerTPK) DeleteCustomer(ctx echo.Context) error {
	customerId, _ := strconv.Atoi(ctx.Param("customerId"))
	if err := c.service.DeleteCustomer(customerId); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	response := map[string]string{
		"message": "success",
	}
	return ctx.JSON(http.StatusOK, response)
}
