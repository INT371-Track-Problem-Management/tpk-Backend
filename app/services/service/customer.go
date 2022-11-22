package service

import (
	"tpk-backend/app/models/model"
	"tpk-backend/app/models/response"
)

func (s serviceTPK) FetchCustomerByEmail(email string) (*model.Customer, error) {
	customer, err := s.repo.CustomerByEmail(email)
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (s serviceTPK) GetListCustomer() ([]*response.ListCustomer, error) {
	customers, err := s.repo.GetListCustomer()
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (s serviceTPK) FetchCustomerById(customerId int) (*model.Customer, error) {
	customer, err := s.repo.GetCustomerById(customerId)
	if err != nil {
		return nil, err
	}
	return customer, nil
}
