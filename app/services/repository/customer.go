package repository

import "tpk-backend/app/models/model"

func (r mysqlRepository) CustomerByEmail(email string) (*model.Customer, error) {
	var Customer model.Customer
	err := r.conn.Table("customer").Where("email = ?", email).Find(&Customer).Error
	if err != nil {
		return nil, err
	}
	return &Customer, nil
}

func (r mysqlRepository) GetCustomerById(customerId int) (*model.Customer, error) {
	var Customer *model.Customer
	err := r.conn.Table("customer").Where("customerId = ?", customerId).Find(&Customer).Error
	if err != nil {
		return nil, err
	}
	return Customer, nil
}