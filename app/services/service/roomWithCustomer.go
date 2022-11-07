package service

import "tpk-backend/app/models/model"

func (s serviceTPK) GetRoomWithCustomerId(customerId int) (*[]model.RoomWithCustomerId, error) {
	rooms, err := s.repo.GetRoomWithCustomerId(customerId)
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

func (s serviceTPK) GetAllRoomWithCustomerByBuildingId(buildingId int) ([]*model.RoomJoinBulding, error) {
	rooms, err := s.repo.GetAllRoomWithCustomerByBuildingId(buildingId)
	if err != nil {
		return nil, err
	}
	return rooms, nil
}
