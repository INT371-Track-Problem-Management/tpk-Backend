package service

import "tpk-backend/app/models/response"

func (s serviceTPK) AllBuilding() (*[]response.AllBuilding, error) {
	buildings, err := s.repo.AllBuilding()
	if err != nil {
		return nil, err
	}
	return buildings, nil
}
