package repository

import (
	"tpk-backend/app/models/model"
	"tpk-backend/app/models/request"
)

func (r mysqlRepository) RegisUser(req model.User) error {
	err := r.conn.Table("userApp").Create(&req).Error
	if err != nil {
		return err
	}
	return nil
}

func (r mysqlRepository) RegisterCustomersRepo(req request.CustomerRegis) (*int, error) {
	var err error
	stmt := r.conn.Begin()
	err = stmt.Table("customer").Create(&req).Error
	if err != nil {
		return nil, err
	}
	var cusid int
	err = stmt.Table("customer").Select("customerId").Where("email = ?", req.Email).Scan(&cusid).Error
	if err != nil {
		return nil, err
	}
	stmt.Commit()
	return &cusid, nil
}

func (r mysqlRepository) RegisterEmployeeRepo(req re.EmployeeRegis) (*int, error) {
	stmt := r.conn.Begin()
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
