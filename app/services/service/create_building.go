package service

import (
	"tpk-backend/app/models/model"
	"tpk-backend/app/models/request"
	"tpk-backend/app/pkg"
)

func (s serviceTPK) CreateBuilding(req request.BuildingInsert) (*int64, error) {
	now := pkg.GetDatetime()
	model := model.BuildingInsert{
		BuildingName: req.BuildingName,
		CreateAt:     now,
		UpdateAt:     now,
		UpdateBy:     req.UpdateBy,
	}
	buildingId, err := s.repo.CreateBuilding(model)
	if err != nil {
		return nil, err
	}
	return buildingId, nil
}
