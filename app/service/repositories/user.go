package repositories

import (
	"fmt"
	"tpk-backend/app/model/entity"
	"tpk-backend/app/model/request"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisUser(ctx echo.Context, conn *gorm.DB, req entity.User) error {
	err := conn.Table("userApp").Create(&req).Error
	if err != nil {
		return err
	}
	return nil
}

func ChangeEmail(ctx echo.Context, conn *gorm.DB, req request.ChangeEmail, oldEmail string) error {
	var err error
	stmt := conn.Begin()
	err = stmt.Table("userApp").Where("email = ?", oldEmail).Update("email = ?", req.NewEmail).Error
	if err != nil {
		stmt.Rollback()
		return err
	}
	stmt.Commit()
	return nil
}

func ChangePassword(ctx echo.Context, conn *gorm.DB, model entity.ChangePassword) error {
	var err error
	sql := fmt.Sprintf(`
		UPDATE userApp
		SET password = '%v'
		WHERE email = '%v'
	`,
		model.NewPassword,
		model.Email)
	fmt.Println(sql)
	stmt := conn.Begin()
	err = stmt.Exec(sql).Error
	if err != nil {
		stmt.Rollback()
		return err
	}
	stmt.Commit()
	return nil
}

func GetProfileEmpByEmail(ctx echo.Context, conn *gorm.DB, email string) (*entity.Employee, error) {
	var emp entity.Employee
	err := conn.Table("employee").Where("email = ?", email).Scan(&emp).Error
	if err != nil {
		return nil, err
	}
	return &emp, nil
}

func GetProfileCustomerByEmail(ctx echo.Context, conn *gorm.DB, email string) (*entity.Customer, error) {
	var cus entity.Customer
	err := conn.Table("customer").Where("email = ?", email).Scan(&cus).Error
	if err != nil {
		return nil, err
	}
	return &cus, nil
}
