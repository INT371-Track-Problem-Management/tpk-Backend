package service

import "tpk-backend/app/models/model"

func (s serviceTPK) FetchRoomByRoomNum(roomnum string) (*model.Room, error) {
	room, err := s.repo.FetchRoomByRoomNum(roomnum)
	if err != nil {
		return nil, err
	}
	return room, nil
}
