package service

import "tpk-backend/app/models/model"

func (s serviceTPK) GetAllRoomAndCustomerByBuildingId(buildingId int) ([]*model.RoomJoinBulding, error) {
	rooms, err := s.repo.GetAllRoomAndCustomerByBuildingId(buildingId)
	if err != nil {
		return nil, err
	}
	return rooms, nil
}
