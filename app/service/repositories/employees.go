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
