package repository

func (r mysqlRepository) CheckHealthy() (*string, error) {
	var data string
	err := r.conn.Raw(`SELECT "Healthy" FROM DUAL`).Scan(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}
