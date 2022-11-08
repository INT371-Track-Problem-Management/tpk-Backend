package repository

func (r mysqlRepository) BuildingDelete(buildingId int) error {
	session := r.conn.Begin()

	if err := session.Exec("DELETE FROM reportEngage WHERE buildingId = ?", buildingId).Error; err != nil {
		return err
	}

	if err := session.Exec("DELETE FROM roomWithCustomer WHERE buildingId = ?", buildingId).Error; err != nil {
		return err
	}

	if err := session.Exec("DELETE FROM room WHERE buildingId = ?", buildingId).Error; err != nil {
		return err
	}

	if err := session.Exec("DELETE FROM building WHERE buildingId = ?", buildingId).Error; err != nil {
		return err
	}

	session.Commit()
	return nil
}
