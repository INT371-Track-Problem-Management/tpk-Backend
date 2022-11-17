package repository

func (r mysqlRepository) DeleteProfileMedia(email string) error {
	if err := r.conn.Exec("DELETE FROM profileMedia WHERE email = ?", email).Error; err != nil {
		return err
	}
	return nil
}
