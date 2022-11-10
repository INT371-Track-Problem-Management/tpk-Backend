package repository

func (r mysqlRepository) DeleteCustomer(id int) error {
	if err := r.conn.Exec("DELETE FROM customer WHERE customerId = ?", id).Error; err != nil {
		return err
	}
	return nil
}
