package service

import "tpk-backend/app/models/model"

func (s serviceTPK) FetchEmployeeByEmail(email string) (*model.Employee, error) {
	employee, err := s.repo.EmployeeByEmail(email)
	if err != nil {
		return nil, err
	}
	return employee, nil
}
