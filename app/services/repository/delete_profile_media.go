package repository

import "gorm.io/gorm"

func (r mysqlRepository) DeleteProfileMedia(email string, session *gorm.DB) error {
	if err := session.Exec("DELETE FROM profileMedia WHERE email = ?", email).Error; err != nil {
		return err
	}
	return nil
}
