package repository

import "tpk-backend/app/models/response"

func (r mysqlRepository) AllBuilding() (*[]response.AllBuilding, error) {
	var building []response.AllBuilding
	err := r.conn.Table("building").Find(&building).Error
	if err != nil {
		return nil, err
	}
	return &building, nil
}
