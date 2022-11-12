package repository

import "tpk-backend/app/models/model"

func (r mysqlRepository) ReportMediaById(id string) (*model.ReportMedia, error) {
	var media model.ReportMedia
	err := r.conn.Table("reportMedia").Where("id = ?", id).Find(&media).Error
	if err != nil {
		return nil, err
	}
	return &media, nil
}
