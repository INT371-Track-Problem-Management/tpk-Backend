package repository

func (repo mysqlRepository) CheckHealthy() (*string, error) {
	var data string
	err := repo.conn.Raw(`SELECT "Healthy" FROM DUAL`).Scan(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}
