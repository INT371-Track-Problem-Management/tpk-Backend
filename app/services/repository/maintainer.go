package repository

import "tpk-backend/app/models/model"

func (r mysqlRepository) Maintainerlist() ([]*model.Maintainer, error) {
	var maintainers []*model.Maintainer
	err := r.conn.Table("maintainer").Find(&maintainers).Error
	if err != nil {
		return nil, err
	}
	return maintainers, nil
}

func (r mysqlRepository) MaintainerById(maintainerId int) (*model.Maintainer, error) {
	var maintainers model.Maintainer
	err := r.conn.Table("maintainer").Where("maintainerId = ?", maintainerId).Scan(&maintainers).Error
	if err != nil {
		return nil, err
	}
	return &maintainers, nil
}
