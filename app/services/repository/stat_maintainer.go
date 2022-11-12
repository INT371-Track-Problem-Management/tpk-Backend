package repository

import "tpk-backend/app/models/model"

func (r mysqlRepository) FetchStatMaintain() (*model.StatMaintainer, error) {
	var stat model.StatMaintainer
	sql := `
	select
		m.maintainerId,
		m.fname,
		m.lname,
	AVG(rr.score) as average
	from
		maintainer m
	left join
		reviewReports rr 
	on
		m.maintainerId = rr.maintainertId 
	group by
		m.maintainerId `
	if err := r.conn.Raw(sql).Scan(&stat).Error; err != nil {
		return nil, err
	}
	return &stat, nil
}
