package repositories

import (
	"fmt"
	entity "tpk-backend/app/model/entity"
	"tpk-backend/app/model/request"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RoomAddCustomer(ctx echo.Context, conn *gorm.DB, req request.RoomAddCustomer) error {
	var err error
	stmt := conn.Begin()

	err = stmt.Exec(`
	INSERT INTO roomWithCustomer (roomId, customerId, status, dormId)
	VALUES (?, ?, ?, ?)
	`,
		req.RoomId,
		req.CustomerId,
		"A",
		req.DoomId).Error
	if err != nil {
		stmt.Rollback()
		return err
	}

	err = stmt.Exec("UPDATE room SET status = ? WHERE roomId = ?", "A", req.RoomId).Error
	if err != nil {
		stmt.Rollback()
		return err
	}

	stmt.Commit()
	return nil
}

func RoomRemoveCustomer(ctx echo.Context, conn *gorm.DB, id int) error {
	var err error
	var rwc entity.RoomWithCustomer
	stmt := conn.Begin()
	err = stmt.Table("roomWithCustomer").Where("id = ?", id).Find(&rwc).Error
	if err != nil {
		stmt.Rollback()
		return err
	}

	err = stmt.Exec("UPDATE roomWithCustomer SET status = ? WHERE id = ?", "I", rwc.Id).Error
	if err != nil {
		stmt.Rollback()
		return err
	}

	err = stmt.Exec("UPDATE room SET status = ? WHERE roomId = ?", "A", rwc.RoomId).Error
	if err != nil {
		stmt.Rollback()
		return err
	}

	stmt.Commit()
	return nil
}

func GetAllRoomWithCustomer(ctx echo.Context, conn *gorm.DB, buildingId int) ([]*entity.RoomJoinDorm, error) {
	var result []*entity.RoomJoinDorm

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
		r.buildingId as buildingId	
	FROM
		roomWithCustomer rwc
	JOIN room r
	ON
		r.roomId = rwc.roomId
	WHERE
		r.buildingId = %v
		`, buildingId)

	err := conn.Raw(sql).Scan(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetRoomWithCustomerByCustomerId(ctx echo.Context, conn *gorm.DB, customerId int) (*entity.RoomWithCustomer, error) {
	roomProfile := new(entity.RoomWithCustomer)

	err := conn.Table("roomWithCustomer").Where("customerId = ?", customerId).Find(roomProfile).Error
	if err != nil {
		return nil, err
	}
	return roomProfile, nil
}
