package service

import (
	"tpk-backend/app/models/model"
	"tpk-backend/app/models/request"
)

func (s serviceTPK) ChangePassword(req request.ChangePassword) error {
	model := model.ChangePassword{
		Email:       req.Email,
		OldPassword: req.OldPassword,
		NewPassword: req.NewPassword,
	}
	if err := s.repo.ChangePassword(model); err != nil {
		return err
	}
	return nil
}
