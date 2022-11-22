package service

import (
	"fmt"
	"tpk-backend/app/constants"
	"tpk-backend/app/jwt"
	"tpk-backend/app/models/model"
	"tpk-backend/app/pkg"
)

func (s serviceTPK) ForgetPassword(user model.User) error {
	jwt := jwt.DecodeJWTPassword(user.Password)
	body := fmt.Sprintf(`เรียนผู้ใช้อีเมลล์ %v เราได้รับคำร้องการลืมรหัสผ่าน รหัสผ่านของคุณคือ: %v`, user.Email, jwt.Password)

	if err := pkg.Smtp2(constants.SUBJECT_EMAIL_FORGET_PASSWORD, user.Email, body); err != nil {
		return err
	}

	return nil
}
