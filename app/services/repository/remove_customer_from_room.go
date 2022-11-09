package repository

import "tpk-backend/app/models/model"

func (r mysqlRepository) RemoveCustomerFromRoom(id int, now string) error {
	var err error
	var rwc model.RoomWithCustomer
	stmt := r.conn.Begin()
	err = stmt.Table("roomWithCustomer").Where("id = ?", id).Find(&rwc).Error
	if err != nil {
		stmt.Rollback()
		return err
	}

	err = stmt.Exec("UPDATE roomWithCustomer SET status = ?, updateAt = ? WHERE id = ?", "I", now, rwc.Id).Error
	if err != nil {
		stmt.Rollback()
		return err
	}

	err = stmt.Exec("UPDATE room SET status = ?, updateAt = ? WHERE roomId = ?", "I", now, rwc.RoomId).Error
	if err != nil {
		stmt.Rollback()
		return err
	}

	stmt.Commit()
	return nil
}
