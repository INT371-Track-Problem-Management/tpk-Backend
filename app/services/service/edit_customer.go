package service

import (
	"tpk-backend/app/models/model"
	"tpk-backend/app/models/request"
	"tpk-backend/app/pkg"
)

func (s serviceTPK) CustomerEditProfile(req request.CustomerEditProfile, email string) error {
	timenow := pkg.GetDatetime()
	model := model.CustomerEditProfile{
		Fname:       req.Fname,
		Lname:       req.Lname,
		Sex:         req.Sex,
		DateOfBirth: req.DateOfBirth,
		Age:         req.Age,
		Phone:       req.Phone,
		Address:     req.Address,
		UpdateAt:    timenow,
		UpdateBy:    req.UpdateBy,
	}
	err := s.repo.CustomerEditProfile(model, email)
	if err != nil {
		return err
	}
	return nil
}
