package repository

import "tpk-backend/app/models/model"

func (r mysqlRepository) FetchRoomByRoomNum(roomnum string) (*model.Room, error) {
	var room model.Room
	if err := r.conn.Table("room").Where("roomNum = ?", roomnum).Scan(&room).Error; err != nil {
		return nil, err
	}
	return &room, nil
}
