package service

import (
	"errors"
	"log"
	"tpk-backend/app/jwt"
	"tpk-backend/app/models/request"
	"tpk-backend/app/models/response"
)

func (s serviceTPK) Login(req request.User) (*response.Token, error) {
	user, err := s.repo.GetUser(req.Email)
	if err != nil {
		return nil, err
	}

	if req.Email != user.Email {
		errUn := errors.New("Unatutherize")
		return nil, errUn
	}

	if pwd := jwt.ComparePassword(req.Password, user.Password); !pwd {
		errUn := errors.New("Unatutherize")
		return nil, errUn
	}

	if user.Role == "C" {
		cus, err := s.repo.CustomerByEmail(user.Email)
		if err != nil {
			return nil, err
		}
		token, err := jwt.GenerateTokenLogin(cus.CustomerId, user.Email, user.Role)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		if err := s.repo.SaveToken(token, user.Role); err != nil {
			return nil, err
		}
		res := response.Token{
			Token: *token,
			Name:  cus.Fname,
		}
		return &res, nil
	}

	if user.Role == "E" {
		emp, err := s.repo.EmployeeByEmail(user.Email)
		if err != nil {
			return nil, err
		}
		token, err := jwt.GenerateTokenLogin(emp.EmployeeId, user.Email, user.Role)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		if err := s.repo.SaveToken(token, user.Role); err != nil {
			return nil, err
		}

		res := response.Token{
			Token: *token,
			Name:  emp.Fname,
		}
		return &res, nil
	}

	if user.Role == "A" {
		emp, err := s.repo.EmployeeByEmail(user.Email)
		if err != nil {
			return nil, err
		}
		token, err := jwt.GenerateTokenLogin(emp.EmployeeId, user.Email, user.Role)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		if err := s.repo.SaveToken(token, user.Role); err != nil {
			return nil, err
		}

		res := response.Token{
			Token: *token,
			Name:  emp.Fname,
		}
		return &res, nil
	}

	return nil, nil
}
