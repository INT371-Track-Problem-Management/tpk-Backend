package repository

import "tpk-backend/app/models/model"

func (r mysqlRepository) EditProfile(model model.EditProfile, email string, role string) error {
	if err := r.conn.Table(role).Where("email = ?", email).Updates(model).Error; err != nil {
		return err
	}
	return nil
}
