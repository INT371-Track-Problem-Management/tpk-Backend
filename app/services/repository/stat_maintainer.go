package repository

import "tpk-backend/app/models/model"

func (r mysqlRepository) FetchStatMaintain() (*[]model.StatMaintainer, error) {
	var stat *[]model.StatMaintainer
	sql := `
	select
		m.maintainerId,
		m.fname,
		m.lname,
		m.phone,
		COUNT(rr.reviewId) as total,
		AVG(rr.score) as average
	from
		maintainer m
	left join
		reviewReports rr 
	on
		m.maintainerId = rr.maintainerId 
	group by
		m.maintainerId `
	if err := r.conn.Raw(sql).Scan(&stat).Error; err != nil {
		return nil, err
	}
	return stat, nil
}

func (r mysqlRepository) FetchOverviewMaintain(id int) (*[]model.OverviewMaintainer, error) {
	var stat *[]model.OverviewMaintainer
	sql := `
	select
		rr.score,
		rr.des
	from
		reviewReports rr
	where
		rr.maintainerId = ?`
	if err := r.conn.Raw(sql, id).Scan(&stat).Error; err != nil {
		return nil, err
	}
	return stat, nil
}
