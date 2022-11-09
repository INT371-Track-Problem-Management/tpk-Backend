package service

import (
	"tpk-backend/app/models/model"
	"tpk-backend/app/models/request"
	"tpk-backend/app/pkg"
)

func (s serviceTPK) EditProfile(req request.EditProfile, email string, role string) error {
	timenow := pkg.GetDatetime()
	model := model.EditProfile{
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
	err := s.repo.EditProfile(model, email, role)
	if err != nil {
		return err
	}
	return nil
}
