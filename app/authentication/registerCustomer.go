package authentication

import (
	"fmt"
	"tpk-backend/app/model/entity"
	"tpk-backend/app/model/request"
	"tpk-backend/app/pkg"
	"tpk-backend/app/pkg/config"
	"tpk-backend/app/service/repositories"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type RegisterCustomer struct {
	Password    string `gorm:"column:password" json:"password"`
	Email       string `json:"email" gorm:"column:email"`
	Fname       string `json:"fname" gorm:"column:fname"`
	Lname       string `json:"lname" gorm:"column:lname"`
	Sex         string `json:"sex" gorm:"column:sex"`
	DateOfBirth string `json:"dateOfBirth" gorm:"column:dateOfBirth"`
	Age         int    `json:"age" gorm:"column:age"`
	Phone       string `json:"phone" gorm:"column:phone"`
	Address     string `json:"address" gorm:"column:address"`
}

func RegisterCustomers(ctx echo.Context, conn *gorm.DB, req RegisterCustomer, uri string) (*int, error) {
	cusId, err := RegisterCustomersService(ctx, conn, req, uri)
	if err != nil {
		return nil, err
	}
	return cusId, nil
}

func ActivateCustomerCtr(ctx echo.Context, conn *gorm.DB, tokeni string, status string) error {
	errt := ActivateCustomer(conn, tokeni, status)
	if errt != nil {
		return errt
	}
	return nil
}

func RegisterCustomersService(ctx echo.Context, conn *gorm.DB, req RegisterCustomer, uri string) (*int, error) {

	regisUser := entity.User{
		Email:    req.Email,
		Password: req.Password,
		Role:     "C",
	}
	err := repositories.RegisUser(ctx, conn, regisUser)
	if err != nil {
		return nil, err
	}
	fmt.Println("Register email: " + regisUser.Email + " as a customer")

	regisCus := request.CustomerRegis{
		Email:       req.Email,
		Fname:       req.Fname,
		Lname:       req.Lname,
		Sex:         req.Sex,
		DateOfBirth: req.DateOfBirth,
		Age:         req.Age,
		Phone:       req.Phone,
		Address:     req.Address,
		Status:      "I",
	}
	id, err := RegisterCustomersRepo(ctx, conn, regisCus)
	if err != nil {
		return nil, err
	}
	fmt.Println("Register customer success id is " + fmt.Sprintln(*id))
	rsp := config.LoadRegisCustomerSend()
	if regisCus.Email != "" {
		// token, err := GenerateTokenRegister(*id)
		// if err != nil {
		// 	return nil, err
		// }
		act := fmt.Sprintln(*id)
		activate := uri + "api/activateCus?cusid=" + act
		pkg.SSLemail(&regisCus.Email, rsp.Subject, activate)
	}

	return id, err
}

func ActivateCustomer(conn *gorm.DB, cusId string, status string) error {
	err := ChangeStatus(conn, cusId, status)
	if err != nil {
		return err
	}
	return nil
}

func RegisterCustomersRepo(ctx echo.Context, conn *gorm.DB, req request.CustomerRegis) (*int, error) {
	var err error
	err = conn.Table("customer").Create(&req).Error
	if err != nil {
		fmt.Println("Register customer unsuccess" + err.Error())
		return nil, err
	}
	fmt.Println("Register customer success")
	var cusid int
	err = conn.Table("customer").Select("customerId").Where("email = ?", req.Email).Scan(&cusid).Error
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return &cusid, nil
}

func ChangeStatus(conn *gorm.DB, cusId string, status string) error {
	err := conn.Exec(`UPDATE customer SET status = ? WHERE customerId = ?`, status, cusId).Error
	if err != nil {
		return err
	}
	fmt.Println("status customerId " + cusId + " change to " + status)
	return nil
}
