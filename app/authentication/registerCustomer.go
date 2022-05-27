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
	Username    string `gorm:"column:username" json:"username"`
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

func RegisterCustomers(ctx echo.Context, conn *gorm.DB, req RegisterCustomer) (*int, error) {
	cusId, err := RegisterCustomersService(ctx, conn, req)
	if err != nil {
		return nil, err
	}
	return cusId, nil
}

func RegisterCustomersService(ctx echo.Context, conn *gorm.DB, req RegisterCustomer) (*int, error) {

	regisUser := entity.User{
		Username: req.Username,
		Password: req.Password,
		Role:     "C",
	}
	err := repositories.RegisUser(ctx, conn, regisUser)
	if err != nil {
		return nil, err
	}
	fmt.Println("Register username: " + regisUser.Username + " as a customer")

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
		Username:    req.Username,
	}
	id, err := RegisterCustomersRepo(ctx, conn, regisCus)
	if err != nil {
		return nil, err
	}
	fmt.Println("Register customer success id is " + *id)
	rsp := config.LoadRegisCustomerSend()

	if regisCus.Email != "" {
		pkg.SSLemail(&regisCus.Email, rsp.Subject, "")
	}
	cusId, err := repositories.CustomerByUsername(ctx, conn, req.Username)
	if err != nil {
		return nil, err
	}

	return &cusId.CustomerId, err
}

func RegisterCustomersRepo(ctx echo.Context, conn *gorm.DB, req request.CustomerRegis) (*string, error) {
	var err error
	err = conn.Table("customer").Create(&req).Error
	if err != nil {
		fmt.Println("Register customer unsuccess" + err.Error())
		return nil, err
	}
	fmt.Println("Register customer success")
	var cusid string
	err = conn.Table("customer").Select("customerId").Where("username = ?", req.Username).Scan(&cusid).Error
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return &cusid, nil
}

func ChangeStatus(conn *gorm.DB, cusId string, status string) error {
	err := conn.Table("customer").Update("status = ?", status).Where("customerId = ?", cusId).Error
	if err != nil {
		return err
	}
	fmt.Println("status customerId " + cusId + " change to " + status)
	return nil
}
