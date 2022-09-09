package repositories

import (
	"tpk-backend/app/model/entity"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func EmployeeByEmail(ctx echo.Context, conn *gorm.DB, email string) (*entity.Employee, error) {
	var emp entity.Employee
	err := conn.Table("employee").Where("email = ?", email).Find(&emp).Error
	if err != nil {
		return nil, err
	}
	return &emp, nil
}

func EmployeeById(ctx echo.Context, conn *gorm.DB, id int) (*entity.Employee, error) {
	var emp entity.Employee
	err := conn.Table("employee").Where("employeeId = ?", id).Find(&emp).Error
	if err != nil {
		return nil, err
	}
	return &emp, nil
}

func SelectDormIdByEmployeeId(ctx echo.Context, conn *gorm.DB, employeeId int) (*int, error) {
	var dormId int
	err := conn.Table("employeeWithDorm").Select("dormId").Where("employeeId = ?", employeeId).Scan(&dormId).Error
	if err != nil {
		return nil, err
	}
	return &dormId, nil
}
