package repository

import "tpk-backend/app/models/model"

func (r mysqlRepository) RoomByRoomId(roomId int) (*model.Room, error) {
	var room model.Room
	if err := r.conn.Table("room").Where("roomId = ?", roomId).Find(&room).Error; err != nil {
		return nil, err
	}
	return &room, nil
}
