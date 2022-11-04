package repository

import "tpk-backend/app/models/model"

func (r mysqlRepository) GetUser(email string) (*model.User, error) {
	user := new(model.User)
	err := r.conn.Table("userApp").Where("email = ?", email).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r mysqlRepository) SaveToken(token *string) error {
	save := model.SaveToken{
		Token:  *token,
		Status: `A`,
	}
	err := r.conn.Table("tokenApp").Create(save).Error
	if err != nil {
		return err
	}
	return nil
}
