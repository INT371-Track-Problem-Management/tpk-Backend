package repository

import (
	"fmt"
	"tpk-backend/app/models/model"
)

func (r mysqlRepository) TotalFlooorsByBuildingId(buildingId int) (*int, error) {
	var floors int
	sql := fmt.Sprintf(`
		SELECT MAX(r.floors)
		FROM room r 
		WHERE r.buildingId = %v;
	`, buildingId)
	err := r.conn.Raw(sql).Scan(&floors).Error
	if err != nil {
		return nil, err
	}
	return &floors, nil
}

func (r mysqlRepository) RoomInFloorByBuildingId(buildingId int, floor int) (*[]model.RoomByFloors, error) {
	var rooms []model.RoomByFloors
	sql := fmt.Sprintf(`
		SELECT r.*
		FROM room r 
		WHERE r.buildingId = %v
		AND r.floors = %v;
	`, buildingId, floor)
	err := r.conn.Raw(sql).Scan(&rooms).Error
	if err != nil {
		return nil, err
	}
	return &rooms, nil
}
