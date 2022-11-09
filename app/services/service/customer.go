package service

import "tpk-backend/app/models/model"

func (s serviceTPK) FetchCustomerByEmail(email string) (*model.Customer, error) {
	customer, err := s.repo.CustomerByEmail(email)
	if err != nil {
		return nil, err
	}
	return customer, nil
}
