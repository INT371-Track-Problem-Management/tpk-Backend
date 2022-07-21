package server

import (
	"fmt"
	"net/http"
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
		return ctx.JSON(http.StatusBadRequest, "")
	}
	token, err = authentication.Login(ctx, h.DB, *user)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, err.Error())
	}
	return ctx.JSON(http.StatusOK, echo.Map{"token": token})
}

func (h *FuncHandler) RegisterCustomer(ctx echo.Context) error {
	var err error
	req := new(authentication.RegisterCustomer)
	err = ctx.Bind(&req)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, "")
	}
	customerId, err := authentication.RegisterCustomers(ctx, h.DB, *req, URI)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, "")
	}
	return ctx.JSON(http.StatusOK, customerId)
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
		return ctx.JSON(http.StatusBadRequest, "")
	}
	res, err := controller.TestController(ctx, *req, h.DB)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, "")
	}

	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) CheckHealthy(ctx echo.Context) error {
	res, err := controller.CheckHealthy(ctx, h.DB)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, "")
	}

	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) CheckHealthyJWT(ctx echo.Context) error {
	res, err := controller.CheckHealthy(ctx, h.DB)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, "")
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) GetRoleJWT(ctx echo.Context) error {
	check, status := authentication.ValidateOwnerService(ctx)
	if status == false {
		return ctx.String(http.StatusUnauthorized, check)
	}
	user := authentication.DecodeJWT(ctx)
	return ctx.String(http.StatusOK, user.Role)
}

// Customer service

func (h *FuncHandler) ReportInsert(ctx echo.Context) error {
	check, status := authentication.ValidateCustomerService(ctx)
	if status == false {
		return ctx.String(http.StatusUnauthorized, check)
	}
	req := new(request.ReportInsert)
	err := ctx.Bind(&req)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, "")
	}
	res, err := controller.ReportInsert(ctx, h.DB, *req)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, "")
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) ActivateCustomer(ctx echo.Context) error {
	check, status := authentication.ValidateCustomerService(ctx)
	if status == false {
		return ctx.String(http.StatusUnauthorized, check)
	}
	id := ctx.QueryParam("cusid")
	err := authentication.ActivateCustomerCtr(ctx, h.DB, id, "A")
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, "")
	}
	redir := URI + "login"
	fmt.Println("----test----")
	fmt.Println(redir)
	return ctx.Redirect(http.StatusMovedPermanently, redir)
}

func (h *FuncHandler) GetReportByCreatedBy(ctx echo.Context) error {
	check, status := authentication.ValidateCustomerService(ctx)
	if status == false {
		return ctx.String(http.StatusUnauthorized, check)
	}
	req := new(request.ReportByCreatedBy)
	err := ctx.Bind(&req)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, "")
	}
	res, err := controller.GetReportByCreatedBy(ctx, h.DB, *req)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, "")
	}
	return ctx.JSON(http.StatusOK, res)
}

// Owner service

func (h *FuncHandler) Rooms(ctx echo.Context) error {
	check, status := authentication.ValidateOwnerService(ctx)
	if status == false {
		return ctx.String(http.StatusUnauthorized, check)
	}
	res, err := controller.Rooms(ctx, h.DB)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, "")
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) RoomsStatus(ctx echo.Context) error {
	check, status := authentication.ValidateOwnerService(ctx)
	if status == false {
		return ctx.String(http.StatusUnauthorized, check)
	}
	req := new(request.RoomsStatus)
	err := ctx.Bind(&req)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, "")
	}
	res, err := controller.RoomsStatus(ctx, h.DB, *req)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, "")
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) Customer(ctx echo.Context) error {
	check, status := authentication.ValidateOwnerService(ctx)
	if status == false {
		return ctx.String(http.StatusUnauthorized, check)
	}
	res, err := controller.Customer(ctx, h.DB)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, "")
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) Dorm(ctx echo.Context) error {
	check, status := authentication.ValidateOwnerService(ctx)
	if status == false {
		return ctx.String(http.StatusUnauthorized, check)
	}
	req := new(request.Dorm)
	err := ctx.Bind(&req)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, "")
	}
	res, err := controller.Dorm(ctx, h.DB, *req)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, "")
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) Report(ctx echo.Context) error {
	check, status := authentication.ValidateOwnerService(ctx)
	if status == false {
		return ctx.String(http.StatusUnauthorized, check)
	}
	res, err := controller.Report(ctx, h.DB)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, "")
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) ReportById(ctx echo.Context) error {
	check, status := authentication.ValidateOwnerService(ctx)
	if status == false {
		return ctx.String(http.StatusUnauthorized, check)
	}
	req := new(request.Report)
	err := ctx.Bind(&req)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, "")
	}
	res, err := controller.ReportById(ctx, h.DB, *req)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, "")
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) DormInsert(ctx echo.Context) error {
	check, status := authentication.ValidateOwnerService(ctx)
	if status == false {
		return ctx.String(http.StatusUnauthorized, check)
	}
	req := new(request.DormInsert)
	err := ctx.Bind(&req)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, "")
	}
	res, err := controller.DormInsert(ctx, h.DB, *req)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, "")
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) RoomsInsert(ctx echo.Context) error {
	check, status := authentication.ValidateOwnerService(ctx)
	if status == false {
		return ctx.String(http.StatusUnauthorized, check)
	}
	req := new(request.RoomInsert)
	err := ctx.Bind(&req)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, "")
	}
	res, err := controller.RoomInsert(ctx, h.DB, *req)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, "")
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) DormDelete(ctx echo.Context) error {
	check, status := authentication.ValidateOwnerService(ctx)
	if status == false {
		return ctx.String(http.StatusUnauthorized, check)
	}
	req := new(request.DormDelete)
	err := ctx.Bind(&req)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, "")
	}
	res, err := controller.DormDelete(ctx, h.DB, *req)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, "")
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) ReportChangeStatus(ctx echo.Context) error {
	check, status := authentication.ValidateOwnerService(ctx)
	if status == false {
		return ctx.String(http.StatusUnauthorized, check)
	}
	req := new(request.ReportChangeStatus)
	err := ctx.Bind(&req)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, "")
	}
	res, err := controller.ReportChangeStatus(ctx, h.DB, *req)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, "")
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) GetReportEngageAll(ctx echo.Context) error {
	check, status := authentication.ValidateOwnerService(ctx)
	if status == false {
		return ctx.String(http.StatusUnauthorized, check)
	}
	res, err := controller.GetReportEngageAll(ctx, h.DB)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, "")
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) GetReportEngageById(ctx echo.Context) error {
	check, status := authentication.ValidateOwnerService(ctx)
	if status == false {
		return ctx.String(http.StatusUnauthorized, check)
	}
	req := new(request.ReportEngageById)
	err := ctx.Bind(&req)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, "")
	}
	res, err := controller.GetReportEngageById(ctx, h.DB, *req)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, "")
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) InsertReportEngage(ctx echo.Context) error {
	check, status := authentication.ValidateOwnerService(ctx)
	if status == false {
		return ctx.String(http.StatusUnauthorized, check)
	}
	req := new(request.ReportEngage)
	err := ctx.Bind(&req)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, "")
	}
	res, err := controller.InsertReportEngage(ctx, h.DB, *req)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, "")
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *FuncHandler) DeleteReportById(ctx echo.Context) error {
	check, status := authentication.ValidateOwnerService(ctx)
	if status == false {
		return ctx.String(http.StatusUnauthorized, check)
	}
	var err error
	req := new(request.Report)
	err = ctx.Bind(&req)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, "")
	}
	err = controller.DeleteReportById(ctx, h.DB, *req)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, "")
	}
	return ctx.JSON(http.StatusNoContent, "")
}
