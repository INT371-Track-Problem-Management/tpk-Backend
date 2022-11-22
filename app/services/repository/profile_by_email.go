package repository

import "tpk-backend/app/models/model"

func (r mysqlRepository) ProfileMediaByEmail(email string) (*model.ReportMedia, error) {
	var media model.ReportMedia
	if err := r.conn.Table("profileMedia").Where("email = ?", email).Find(&media).Error; err != nil {
		return nil, err
	}
	return &media, nil
}
