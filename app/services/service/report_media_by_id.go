package service

import "tpk-backend/app/models/model"

func (s serviceTPK) ReportMediaById(id string) (*model.ReportMedia, error) {
	image, err := s.repo.ReportMediaById(id)
	if err != nil {
		return nil, err
	}
	return image, nil
}
