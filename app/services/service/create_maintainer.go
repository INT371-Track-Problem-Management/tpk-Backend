package service

import (
	"tpk-backend/app/models/model"
	"tpk-backend/app/models/request"
	"tpk-backend/app/pkg"
)

func (s serviceTPK) CreateMaintainer(req request.Maintainer) error {
	now := pkg.GetDatetime()
	model := model.AddMaintainer{
		Fname:    req.Fname,
		Lname:    req.Lname,
		Phone:    req.Phone,
		CreateAt: now,
		UpdateAt: now,
		UpdateBy: req.UpdateBy,
	}
	if err := s.repo.CreateMaintainer(model); err != nil {
		return err
	}
	return nil
}
