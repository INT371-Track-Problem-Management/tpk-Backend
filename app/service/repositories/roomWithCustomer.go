package repositories

import (
	"errors"
	"fmt"
	entity "tpk-backend/app/model/entity"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RoomAddCustomer(ctx echo.Context, conn *gorm.DB, model entity.RoomAddCustomer) error {
	var err error
	roomId := fmt.Sprintf(`%v`, model.RoomId)
	checkStatus, err := RoomByRoomId(ctx, conn, roomId)
	if err != nil {
		return err
	}
	if checkStatus.Status == "A" {
		return errors.New("room taken")
	}
	stmt := conn.Begin()
	err = stmt.Table("roomWithCustomer").Create(&model).Error
	if err != nil {
		stmt.Rollback()
		return err
	}
	err = stmt.Exec("UPDATE room SET status = ?, updateAt = ?, updateBy = ? WHERE roomId = ?", "A", model.UpdateAt, model.UpdateBy, model.RoomId).Error
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

func GetAllRoomWithCustomer(ctx echo.Context, conn *gorm.DB, buildingId int) ([]*entity.RoomJoinBulding, error) {
	var result []*entity.RoomJoinBulding

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

	err := conn.Raw(sql).Scan(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetRoomWithCustomerId(ctx echo.Context, conn *gorm.DB, customerId string) (*[]entity.RoomWithCustomerId, error) {
	rooms := new([]entity.RoomWithCustomerId)
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
	err := conn.Raw(sql).Scan(rooms).Error
	if err != nil {
		return nil, err
	}
	return rooms, nil
}
