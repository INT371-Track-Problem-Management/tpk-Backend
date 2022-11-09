package service

import "tpk-backend/app/models/model"

func (s serviceTPK) Maintainerlist() ([]*model.Maintainer, error) {
	maintainers, err := s.repo.Maintainerlist()
	if err != nil {
		return nil, err
	}
	return maintainers, nil
}

func (s serviceTPK) MaintainerById(maintainerId int) (*model.Maintainer, error) {
	maintainer, err := s.repo.MaintainerById(maintainerId)
	if err != nil {
		return nil, err
	}
	return maintainer, nil
}
