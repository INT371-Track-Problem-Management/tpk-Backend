package model

type User struct {
	Email    string `gorm:"column:email" json:"email"`
	Password string `gorm:"column:password" json:"password"`
	Role     string `gorm:"column:role" json:"role"`
}

type ChangePassword struct {
	Email       string `gorm:"column:email" json:"email"`
	OldPassword string `gorm:"column:password" json:"oldPassword"`
	NewPassword string `gorm:"column:password" json:"newPassword"`
}
