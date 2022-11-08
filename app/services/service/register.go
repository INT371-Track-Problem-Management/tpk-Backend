package service

import (
	"tpk-backend/app/jwt"
	"tpk-backend/app/models/model"
	"tpk-backend/app/models/request"
	"tpk-backend/app/pkg"
)

func (s serviceTPK) RegisterCustomersService(req request.RegisterCustomer) (*int, error) {
	var err error
	encryp, err := jwt.GenerateTokenFromPassword(req.Password)
	if err != nil {
		return nil, err
	}
	regisUser := model.User{
		Email:    req.Email,
		Password: *encryp,
		Role:     "C",
	}
	err = s.repo.RegisUser(regisUser)
	if err != nil {
		return nil, err
	}
	timenow := pkg.GetDatetime()
	regisCus := request.CustomerRegis{
		Email:       req.Email,
		Fname:       req.Fname,
		Lname:       req.Lname,
		Sex:         req.Sex,
		DateOfBirth: req.DateOfBirth,
		Age:         req.Age,
		Phone:       req.Phone,
		Address:     req.Address,
		Status:      "I",
		CreateAt:    timenow,
	}
	id, err := s.repo.RegisterCustomersRepo(regisCus)
	if err != nil {
		return nil, err
	}
	return id, err
}

func (s serviceTPK) RegisterOwnerService(req request.RegisterOwner) (*int, error) {
	var err error
	encryp, err := jwt.GenerateTokenFromPassword(req.Password)
	if err != nil {
		return nil, err
	}
	regisUser := model.User{
		Email:    req.Email,
		Password: *encryp,
		Role:     "E",
	}
	err = s.repo.RegisUser(regisUser)
	if err != nil {
		return nil, err
	}

	timenow := pkg.GetDatetime()
	regis := model.EmployeeRegis{
		Fname:       req.Fname,
		Lname:       req.Lname,
		Phone:       req.Phone,
		Address:     req.Address,
		Sex:         req.Sex,
		Email:       req.Email,
		Age:         req.Age,
		DateOfBirth: req.DateOfBirth,
		Position:    req.Position,
		CreateAt:    timenow,
	}

	empId, err := s.repo.RegisterEmployeeRepo(regis)
	if err != nil {
		return nil, err
	}

	return empId, nil
}
