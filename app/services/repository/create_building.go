package repository

import "tpk-backend/app/models/model"

func (r mysqlRepository) CreateBuilding(model model.BuildingInsert) (*int64, error) {
	err := r.conn.Table("building").Create(&model).Error
	if err != nil {
		return nil, err
	}
	var id int64
	err = r.conn.Table("building").Where("createAt = ?", model.CreateAt).Select("buildingId").Scan(&id).Error
	if err != nil {
		return nil, err
	}
	return &id, nil
}
