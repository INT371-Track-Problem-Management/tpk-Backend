package service

import (
	"tpk-backend/app/models/model"
	"tpk-backend/app/models/request"
)

func (s serviceTPK) FetcStatDashBoard(req request.Stat) (*model.Stat, error) {
	stat, err := s.repo.FetcStatDashBoard(req)
	if err != nil {
		return nil, err
	}
	return stat, nil
}
