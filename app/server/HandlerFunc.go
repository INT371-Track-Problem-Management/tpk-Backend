package server

import (
	"fmt"
	"log"
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
	var err error
	user := new(request.User)
	err = ctx.Bind(&user)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	authen, err := authentication.Login(ctx, h.DB, *user)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, err.Error())
	}
	response := map[string]string{
		"token": authen.Token,
		"name":  authen.Name,
	}
	return ctx.JSON(http.StatusOK, response)
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

func (h *FuncHandler) TestSMTP2(ctx echo.Context) error {
	to := "sun_vijitpanmai@hotmail.com"
	// to := "artid.vijitpanmai@mail.kmutt.ac.th"
	// to := "zayori999@gmail.com"
	// to := "paradios.00riser@gmail.com"
	sub := "Test send mail"
	body := "Hello-World"
	err := pkg.Smtp2(sub, to, body)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
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

func (h *FuncHandler) YearConfig(ctx echo.Context) error {
	res, err := controller.YearConfig(ctx, h.DB)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) ReportByRoomId(ctx echo.Context) error {
	roomId := ctx.Param("roomId")
	log.Println(roomId)
	if roomId == "" {
		msg := map[string]string{
			"message": "Require param roomId",
		}
		return ctx.JSON(http.StatusBadRequest, msg)
	}
	res, err := controller.ReportByRoomId(ctx, h.DB, roomId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) ChangePassword(ctx echo.Context) error {
	req := new(request.ChangePassword)
	err := ctx.Bind(&req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	err = controller.ChangePassword(ctx, h.DB, *req)
	if err != nil {
		errmsg := map[string]interface{}{
			"error": err.Error(),
		}
		return ctx.JSON(http.StatusInternalServerError, errmsg)
	}
	res := map[string]string{
		"message": "success",
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) MaintainerById(ctx echo.Context) error {
	param := ctx.Param("maintainerId")
	res, err := controller.MaintainerById(ctx, h.DB, param)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) FetchProfile(ctx echo.Context) error {
	param := ctx.Param("email")
	profile, err := controller.GetProfileByEmail(ctx, h.DB, param)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	response := map[string]interface{}{
		"Profile": profile,
	}
	return ctx.JSON(http.StatusOK, response)
}

func (h *FuncHandler) FetchReportStatusApplication(ctx echo.Context) error {
	reportId := ctx.Param("reportId")
	statuslist, err := controller.ReportStatusByReportId(ctx, h.DB, reportId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	response := map[string]interface{}{
		"Profile": statuslist,
	}
	return ctx.JSON(http.StatusOK, response)
}

func (h *FuncHandler) ReportChangeStatus(ctx echo.Context) error {
	req := new(request.ReportChangeStatus)
	err := ctx.Bind(&req)
	if err != nil {

		return ctx.JSON(http.StatusInternalServerError, err)
	}
	res, err := controller.ReportChangeStatus(ctx, h.DB, *req)
	if err != nil {

		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, res)
}