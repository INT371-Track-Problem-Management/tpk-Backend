package service

import (
	"fmt"
	"tpk-backend/app/models/response"
)

func (s serviceTPK) RoomByBuildingId(buildingId int) (*response.RoomByBuildingId, error) {
	totalFloor, err := s.repo.TotalFlooorsByBuildingId(buildingId)
	if err != nil {
		return nil, err
	}
	rooms := []interface{}{}
	for i := 1; i < *totalFloor+1; i++ {
		floorNum := fmt.Sprintf(`floor%v`, i)
		allroomfloor, err := s.repo.RoomInFloorByBuildingId(buildingId, i)
		if err != nil {
			return nil, err
		}
		room := map[string]interface{}{
			floorNum: allroomfloor,
		}
		rooms = append(rooms, room)
	}

	allroom := response.RoomByBuildingId{
		BuildingId: buildingId,
		Floors:     rooms,
	}

	return &allroom, nil
}
