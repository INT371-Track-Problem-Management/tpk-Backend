package repositories

import (
	entity "tpk-backend/app/model/entity"
	"tpk-backend/app/model/request"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Rooms(ctx echo.Context, conn *gorm.DB) (*[]entity.Room, error) {
	var room []entity.Room
	err := conn.Table("room").Find(&room).Error
	if err != nil {
		return nil, err
	}
	return &room, nil
}

func RoomsStatus(ctx echo.Context, conn *gorm.DB, req request.RoomsStatus) error {
	err := conn.Exec("UPDATE room SET status = ? WHERE roomId = ?", req.Status, req.RoomId).Error
	if err != nil {
		return err
	}
	return nil
}

func RoomInsert(ctx echo.Context, conn *gorm.DB, req request.RoomInsert) error {
	err := conn.Table("room").Create(&req).Error
	if err != nil {
		return err
	}
	return nil
}

func RoomByDormId(ctx echo.Context, conn *gorm.DB, dormId string) (*[]entity.Room, error) {
	var room []entity.Room
	err := conn.Table("room").Where("dormId = ?", dormId).Find(&room).Error
	if err != nil {
		return nil, err
	}
	return &room, nil
}

func RoomAddCustomer(ctx echo.Context, conn *gorm.DB, req request.RoomAddCustomer) error {
	var err error
	stmt := conn.Begin()

	err = stmt.Exec(`
	INSERT INTO roomWithCustomer (roomId, customerId, status)
	VALUES (?, ?, ?)
	`,
		req.RoomId,
		req.CustomerId,
		"A").Error
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
