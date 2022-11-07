package repository

import (
	"fmt"
	"tpk-backend/app/models/model"
)

func (r mysqlRepository) GetRoomWithCustomerId(customerId int) (*[]model.RoomWithCustomerId, error) {
	rooms := new([]model.RoomWithCustomerId)
	sql := fmt.Sprintf(`
		SELECT 
			rwc.roomId,
			r.roomNum,
			rwc.buildingId,
			r.floors,
			r.status 
		FROM room r 
		LEFT JOIN roomWithCustomer rwc
		ON r.roomId = rwc.roomId
		WHERE rwc.customerId = %v;
	`, customerId)
	err := r.conn.Raw(sql).Scan(rooms).Error
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

func (r mysqlRepository) GetAllRoomWithCustomerByBuildingId(buildingId int) ([]*model.RoomJoinBulding, error) {
	var result []*model.RoomJoinBulding
	sql := fmt.Sprintf(`
	SELECT
		rwc.Id as id,
		rwc.roomId as roomId,
		rwc.customerId as customerId,
		rwc.status as status,
		rwc.createAt as createAt,
		rwc.updateAt as updateAt,
		rwc.updateBy as updateBy,
		r.roomNum as roomNum,
		r.floors as floors,
		r.description as description,
		r.buildingId as buildingId,
		c.fname as fname,
        c.lname as lname
	FROM
		roomWithCustomer rwc
	JOIN room r
	ON
		r.roomId = rwc.roomId
	JOIN customer c
    ON
		rwc.customerId = c.customerId
	WHERE
		r.buildingId = %v
		`, buildingId)

	err := r.conn.Raw(sql).Scan(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}
