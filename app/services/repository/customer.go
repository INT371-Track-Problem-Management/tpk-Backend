package repository

import (
	"fmt"
	"tpk-backend/app/models/model"
	"tpk-backend/app/models/response"
)

func (r mysqlRepository) CustomerByEmail(email string) (*model.Customer, error) {
	var Customer model.Customer
	err := r.conn.Table("customer").Where("email = ?", email).Find(&Customer).Error
	if err != nil {
		return nil, err
	}
	return &Customer, nil
}

func (r mysqlRepository) GetListCustomer() ([]*response.ListCustomer, error) {
	var Customer []*response.ListCustomer
	sql :=
		`
	SELECT
    	c.customerId, c.fname, c.lname, c.email, rwc.roomId, r.roomNum, r.floors, rwc.buildingId
	FROM
    	customer c
	LEFT JOIN
    	roomWithCustomer rwc
	ON
    	rwc.customerId = c.customerId
	LEFT JOIN
    	room r 
	ON
    	r.roomId = rwc.roomId
	`
	if err := r.conn.Raw(sql).Scan(&Customer).Error; err != nil {
		return nil, err
	}
	return Customer, nil
}

func (r mysqlRepository) GetCustomerById(customerId int) (*model.Customer, error) {
	var Customer *model.Customer
	err := r.conn.Table("customer").Where("customerId = ?", customerId).Find(&Customer).Error
	if err != nil {
		return nil, err
	}
	return Customer, nil
}

func (r mysqlRepository) GetEmailCreateByReportId(reportId int) (*string, error) {
	var email string
	sql := fmt.Sprintf(
		`
	SELECT
		c.email 
	FROM
		customer c
	LEFT JOIN
		reports r
	ON
		c.customerId = r.createBy
	WHERE 
		r.reportId = %v
	`, reportId)

	if err := r.conn.Raw(sql).Scan(&email).Error; err != nil {
		return nil, err
	}
	return &email, nil
}
