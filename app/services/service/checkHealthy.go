package service

func (s serviceTPK) CheckHealthy() (*string, error) {
	healthy, err := s.repo.CheckHealthy()
	if err != nil {
		return nil, err
	}
	return healthy, nil
}
