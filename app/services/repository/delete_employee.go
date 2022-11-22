package repository

func (r mysqlRepository) DeleteEmployee(id int) error {
	if err := r.conn.Exec("DELETE FROM employee WHERE employeeId = ?", id).Error; err != nil {
		return err
	}
	return nil
}
