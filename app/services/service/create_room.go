package service

import (
	"tpk-backend/app/models/model"
	"tpk-backend/app/models/request"
	"tpk-backend/app/pkg"
)

func (s serviceTPK) CreateRoom(req request.RoomInsert) error {
	session := s.database.Begin()
	now := pkg.GetDatetime()
	for _, room := range req.Rooms {
		model := model.RoomInsert{
			RoomNum:     room.RoomNum,
			Floors:      room.Floors,
			Description: room.Description,
			BuildingId:  req.BuildingId,
			Status:      "I",
			UpdateAt:    now,
			UpdateBy:    req.UpdateBy,
			CreateAt:    now,
		}
		if err := s.repo.CreateRoom(model, session); err != nil {
			return err
		}
	}
	session.Commit()
	return nil
}
