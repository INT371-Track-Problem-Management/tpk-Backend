package repository

import (
	"fmt"
	"tpk-backend/app/models/model"
)

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
