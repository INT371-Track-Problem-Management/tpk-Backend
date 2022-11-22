package service

import "tpk-backend/app/models/model"

func (s serviceTPK) GetUser(email string) (*model.User, error) {
	user, err := s.repo.GetUser(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}
