package controller

import (
	"net/http"
	"strconv"
	"tpk-backend/app/jwt"

	"github.com/labstack/echo/v4"
)

func (c controllerTPK) FetchEmployeeByEmail(ctx echo.Context) error {
	user := jwt.DecodeJWT(ctx)

	employee, err := c.service.FetchEmployeeByEmail(user.Email)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	response := map[string]interface{}{
		"employee": employee,
	}
	return ctx.JSON(http.StatusOK, response)
}

func (c controllerTPK) GetListEmployee(ctx echo.Context) error {
	employees, err := c.service.GetListEmployee()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	if len(employees) == 0 {
		return ctx.JSON(http.StatusNoContent, "No content")
	}

	response := map[string]interface{}{
		"employees": employees,
	}
	return ctx.JSON(http.StatusOK, response)
}

func (c controllerTPK) FetchEmployeeById(ctx echo.Context) error {
	employeeId, _ := strconv.Atoi(ctx.Param("employeeId"))
	employee, err := c.service.FetchEmployeeById(employeeId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	response := map[string]interface{}{
		"employee": employee,
	}
	return ctx.JSON(http.StatusOK, response)
}

func (c controllerTPK) DeleteEmployee(ctx echo.Context) error {
	employeeId, _ := strconv.Atoi(ctx.Param("employeeId"))
	if err := c.service.DeleteEmployee(employeeId); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	response := map[string]string{
		"message": "success",
	}
	return ctx.JSON(http.StatusOK, response)
}
