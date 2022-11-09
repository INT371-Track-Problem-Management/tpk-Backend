package repository

import "tpk-backend/app/models/model"

func (r mysqlRepository) CustomerEditProfile(model model.CustomerEditProfile, email string) error {
	err := r.conn.Table("customer").Where("email = ?", email).Updates(model).Error
	if err != nil {
		return err
	}
	return nil
}
