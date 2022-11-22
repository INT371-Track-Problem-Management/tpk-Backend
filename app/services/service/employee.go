package service

import "tpk-backend/app/models/model"

func (s serviceTPK) FetchEmployeeByEmail(email string) (*model.Employee, error) {
	employee, err := s.repo.EmployeeByEmail(email)
	if err != nil {
		return nil, err
	}
	return employee, nil
}

func (s serviceTPK) GetListEmployee() ([]*model.Employee, error) {
	employees, err := s.repo.GetListEmployee()
	if err != nil {
		return nil, err
	}
	return employees, nil
}

func (s serviceTPK) FetchEmployeeById(customerId int) (*model.Employee, error) {
	employee, err := s.repo.EmployeeById(customerId)
	if err != nil {
		return nil, err
	}
	return employee, nil
}
