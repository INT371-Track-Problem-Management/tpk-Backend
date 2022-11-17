package service

import (
	"tpk-backend/app/models/model"
	"tpk-backend/app/models/response"
)

func (s serviceTPK) FetchStatMaintain() (*[]model.StatMaintainer, error) {
	stat, err := s.repo.FetchStatMaintain()
	if err != nil {
		return nil, err
	}
	return stat, nil
}

func (s serviceTPK) FetchOverviewMaintain(maintainerId int) (*response.OverviewMaintainer, error) {
	maintainer, err := s.repo.MaintainerById(maintainerId)
	if err != nil {
		return nil, err
	}
	stat, err := s.repo.FetchOverviewMaintain(maintainerId)
	if err != nil {
		return nil, err
	}
	overview := response.OverviewMaintainer{
		MaintainerId: maintainer.MaintainerId,
		Fname:        maintainer.Fname,
		Lname:        maintainer.Lname,
		Phone:        maintainer.Phone,
		Overview:     *stat,
	}
	return &overview, nil
}
