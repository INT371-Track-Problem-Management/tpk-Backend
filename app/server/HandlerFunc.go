package server

import (
	"fmt"
	"net/http"
	"strconv"
	"tpk-backend/app/authentication"
	"tpk-backend/app/model/request"
	"tpk-backend/app/pkg"
	"tpk-backend/app/service/controller"

	"github.com/labstack/echo/v4"
)

// Free Service
func (h *FuncHandler) Login(ctx echo.Context) error {
	var token *string
	var err error
	user := new(request.User)
	err = ctx.Bind(&user)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	token, err = authentication.Login(ctx, h.DB, *user)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, err.Error())
	}
	return ctx.JSON(http.StatusOK, echo.Map{"token": token})
}

func (h *FuncHandler) Logout(ctx echo.Context) error {
	err := authentication.Logout(ctx, h.DB)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, err)
	}
	res := map[string]string{
		"message": "success",
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) RegisterCustomer(ctx echo.Context) error {
	var err error
	req := new(authentication.RegisterCustomer)
	err = ctx.Bind(&req)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	customerId, err := authentication.RegisterCustomers(ctx, h.DB, *req, URI)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, "this email can not use!!!")
	}
	return ctx.JSON(http.StatusOK, customerId)
}

func (h *FuncHandler) RegisterOwner(ctx echo.Context) error {
	var err error
	req := new(authentication.RegisterOwner)
	err = ctx.Bind(&req)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	empId, err := authentication.RegisterOwnerCtr(ctx, h.DB, *req)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusUnauthorized, "this email can not use!!!")
	}
	return ctx.JSON(http.StatusOK, empId)
}

func (h *FuncHandler) ChangeEmail(ctx echo.Context) error {
	var err error
	req := new(request.ChangeEmail)
	err = ctx.Bind(&req)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	token := authentication.DecodeJWT(ctx)

	err = controller.ChangeEmail(ctx, h.DB, *req, token.Email)
	if err != nil {
		return ctx.JSON(http.StatusBadGateway, err)
	}
	return ctx.JSON(http.StatusCreated, nil)
}

func (h *FuncHandler) TestGmail(ctx echo.Context) error {
	testmail := "artid.vijitpanmai@mail.kmutt.ac.th"
	pkg.SSLemail(&testmail, "Hello-World", "Hi")
	return ctx.JSON(http.StatusOK, "send mail success")
}

func (h *FuncHandler) Test(ctx echo.Context) error {
	req := new(request.Test)
	err := ctx.Bind(&req)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	res, err := controller.TestController(ctx, *req, h.DB)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) CheckHealthy(ctx echo.Context) error {
	res, err := controller.CheckHealthy(ctx, h.DB)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) CheckHealthyJWT(ctx echo.Context) error {
	res, err := controller.CheckHealthy(ctx, h.DB)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) GetRoleJWT(ctx echo.Context) error {
	check, status := authentication.ValidateOwnerService(ctx)
	if status == false {
		return ctx.String(http.StatusUnauthorized, check.Token)
	}
	user := authentication.DecodeJWT(ctx)
	return ctx.String(http.StatusOK, user.Role)
}

func (h *FuncHandler) GetHistoryByHistoryId(ctx echo.Context) error {
	param := ctx.QueryParam("historyId")
	id, _ := strconv.ParseInt(param, 10, 32)
	res, err := controller.GetHistoryByHistoryId(ctx, h.DB, id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, res)
}
