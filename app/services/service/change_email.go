package service

import (
	"errors"
	"tpk-backend/app/models/request"
)

func (s serviceTPK) ChangeEmail(req request.ChangeEmail, oldEmail string) error {
	var err error
	user := request.User{
		Email: oldEmail,
	}

	oldpwd, err := s.repo.GetUser(user.Email)
	if err != nil {
		return err
	}

	if req.Password != oldpwd.Password {
		return errors.New("Invalid_Token")
	}

	err = s.repo.ChangeEmail(req, oldEmail)
	if err != nil {
		return err
	}
	return nil
}
