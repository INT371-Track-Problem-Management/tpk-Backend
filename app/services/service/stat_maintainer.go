package service

import "tpk-backend/app/models/model"

func (s serviceTPK) FetchStatMaintain() (*model.StatMaintainer, error) {
	stat, err := s.repo.FetchStatMaintain()
	if err != nil {
		return nil, err
	}
	return stat, nil
}
