package service

import "tpk-backend/app/models/model"

func (s serviceTPK) ProfileMediaByEmail(email string) (*model.ReportMedia, error) {
	profile, err := s.repo.ProfileMediaByEmail(email)
	if err != nil {
		return nil, err
	}
	return profile, nil
}
