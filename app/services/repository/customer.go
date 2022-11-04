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

// func (r mysqlRepository) RegisterCustomersRepo(req request.CustomerRegis) (*int, error) {
// 	var err error
// 	stmt := conn.Begin()
// 	err = stmt.Table("customer").Create(&req).Error
// 	if err != nil {
// 		fmt.Println("Register customer unsuccess" + err.Error())
// 		return nil, err
// 	}
// 	fmt.Println("Register customer success")
// 	var cusid int
// 	err = stmt.Table("customer").Select("customerId").Where("email = ?", req.Email).Scan(&cusid).Error
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return nil, err
// 	}
// 	stmt.Commit()
// 	return &cusid, nil
// }