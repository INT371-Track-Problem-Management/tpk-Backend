package service

import (
	"tpk-backend/app/jwt"
	"tpk-backend/app/models/model"
	"tpk-backend/app/models/request"
)

func (s serviceTPK) ChangePassword(req request.ChangePassword) error {
	newPwd, _ := jwt.GenerateTokenFromPassword(req.NewPassword)
	model := model.ChangePassword{
		Email:       req.Email,
		OldPassword: req.OldPassword,
		NewPassword: *newPwd,
	}
	if err := s.repo.ChangePassword(model); err != nil {
		return err
	}
	return nil
}
