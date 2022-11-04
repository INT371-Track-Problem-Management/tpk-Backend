package service

func (t serviceTPK) CheckHealthy() (*string, error) {
	healthy, err := t.repo.CheckHealthy()
	if err != nil {
		return nil, err
	}
	return healthy, nil
}
