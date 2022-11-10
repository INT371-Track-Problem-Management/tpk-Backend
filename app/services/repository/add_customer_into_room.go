package repository

import (
	"errors"
	"tpk-backend/app/models/model"
)

func (r mysqlRepository) AddCustomerIntoRoom(model model.RoomAddCustomer) error {
	var err error
	checkStatus, err := r.RoomByRoomId(model.RoomId)
	if err != nil {
		return err
	}
	if checkStatus.Status == "A" {
		return errors.New("room taken")
	}
	stmt := r.conn.Begin()
	if err = stmt.Table("roomWithCustomer").Create(&model).Error; err != nil {
		stmt.Rollback()
		return err
	}
	if err = stmt.Exec("UPDATE room SET status = ?, updateAt = ?, updateBy = ? WHERE roomId = ?", "A", model.UpdateAt, model.UpdateBy, model.RoomId).Error; err != nil {
		stmt.Rollback()
		return err
	}
	stmt.Commit()
	return nil
}
