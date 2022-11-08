package repository

import (
	"tpk-backend/app/models/model"
	"tpk-backend/app/models/request"
)

func (r mysqlRepository) GetUser(email string) (*model.User, error) {
	user := new(model.User)
	err := r.conn.Table("userApp").Where("email = ?", email).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r mysqlRepository) SaveToken(token *string, role string) error {
	save := model.SaveToken{
		Token:  *token,
		Status: `A`,
		Role:   role,
	}
	err := r.conn.Table("tokenApp").Create(save).Error
	if err != nil {
		return err
	}
	return nil
}

func (r mysqlRepository) ChangeEmail(req request.ChangeEmail, oldEmail string) error {
	var err error
	stmt := r.conn.Begin()
	err = stmt.Table("userApp").Where("email = ?", oldEmail).Update("email = ?", req.NewEmail).Error
	if err != nil {
		stmt.Rollback()
		return err
	}
	stmt.Commit()
	return nil
}
