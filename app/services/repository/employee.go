package repository

import "tpk-backend/app/models/model"

func (r mysqlRepository) EmployeeById(id int) (*model.Employee, error) {
	var emp model.Employee
	err := r.conn.Table("employee").Where("employeeId = ?", id).Find(&emp).Error
	if err != nil {
		return nil, err
	}
	return &emp, nil
}

func (r mysqlRepository) EmployeeByEmail(email string) (*model.Employee, error) {
	var emp model.Employee
	err := r.conn.Table("employee").Where("email = ?", email).Find(&emp).Error
	if err != nil {
		return nil, err
	}
	return &emp, nil
}
