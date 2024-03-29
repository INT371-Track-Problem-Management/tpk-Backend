package request

type User struct {
	Email    string `gorm:"column:email" json:"email"`
	Password string `gorm:"column:password" json:"password"`
}

type ChangeEmail struct {
	NewEmail string `gorm:"column:email" json:"newEmail"`
	Password string `gorm:"column:password" json:"newPassword"`
}

type ChangePassword struct {
	Email       string `gorm:"column:email" json:"email"`
	OldPassword string `gorm:"column:password" json:"oldPassword"`
	NewPassword string `gorm:"column:password" json:"newPassword"`
}

type ForgetPassword struct {
	Email string `gorm:"column:email" json:"email"`
}
