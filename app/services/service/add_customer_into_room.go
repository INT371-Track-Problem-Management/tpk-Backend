package service

import (
	"tpk-backend/app/models/model"
	"tpk-backend/app/models/request"
	"tpk-backend/app/pkg"
)

func (s serviceTPK) AddCustomerIntoRoom(req request.RoomAddCustomer) error {
	now := pkg.GetDatetime()
	model := model.RoomAddCustomer{
		RoomId:     req.UpdateBy,
		CustomerId: req.CustomerId,
		BuildingId: req.BuildingId,
		Status:     "A",
		CreateAt:   now,
		UpdateAt:   now,
		UpdateBy:   req.UpdateBy,
	}
	if err := s.repo.AddCustomerIntoRoom(model); err != nil {
		return err
	}
	return nil
}
