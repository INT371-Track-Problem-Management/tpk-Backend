package authentication

import (
	"fmt"
	"tpk-backend/app/model/entity"
	"tpk-backend/app/service/repositories"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type RegisterOwner struct {
	Password    string `gorm:"column:password" json:"password"`
	Email       string `json:"email" gorm:"column:email"`
	Fname       string `json:"fname" gorm:"column:fname"`
	Lname       string `json:"lname" gorm:"column:lname"`
	Sex         string `json:"sex" gorm:"column:sex"`
	DateOfBirth string `json:"dateOfBirth" gorm:"column:dateOfBirth"`
	Age         int    `json:"age" gorm:"column:age"`
	Phone       string `json:"phone" gorm:"column:phone"`
	Address     string `json:"address" gorm:"column:address"`
	Position    string `json:"position" gorm:"column:position"`
}

func RegisterOwnerCtr(ctx echo.Context, conn *gorm.DB, req RegisterOwner) (*int, error) {
	empId, err := RegisterOwnerService(ctx, conn, req)
	if err != nil {
		return nil, err
	}
	return empId, nil
}

func RegisterOwnerService(ctx echo.Context, conn *gorm.DB, req RegisterOwner) (*int, error) {
	var err error
	encryp, err := GenerateTokenFromPassword(req.Password)
	if err != nil {
		return nil, err
	}
	regisUser := entity.User{
		Email:    req.Email,
		Password: *encryp,
		Role:     "E",
	}
	err = repositories.RegisUser(ctx, conn, regisUser)
	if err != nil {
		return nil, err
	}
	fmt.Println("Register email: " + regisUser.Email + " as a Employee")

	timenow := getDatetime()
	regis := entity.EmployeeRegis{
		Fname:       req.Fname,
		Lname:       req.Lname,
		Phone:       req.Phone,
		Address:     req.Address,
		Sex:         req.Sex,
		Email:       req.Email,
		Age:         req.Age,
		DateOfBirth: req.DateOfBirth,
		Position:    req.Position,
		CreateAt:    timenow,
	}

	empId, err := RegisterOwnerRepo(ctx, conn, regis)
	if err != nil {
		return nil, err
	}

	return empId, nil
}

func RegisterOwnerRepo(ctx echo.Context, conn *gorm.DB, req entity.EmployeeRegis) (*int, error) {
	stmt := conn.Begin()
	err := stmt.Table("employee").Create(&req).Error
	if err != nil {
		stmt.Rollback()
		return nil, err
	}
	var empId int
	err = stmt.Table("employee").Select("employeeId").Where("email = ?", req.Email).Scan(&empId).Error
	if err != nil {
		stmt.Rollback()
		return nil, err
	}
	stmt.Commit()
	return &empId, nil
}
