package request

type User struct {
	Email    string `gorm:"column:email" json:"email"`
	Password string `gorm:"column:password" json:"password"`
}

type ChangeEmail struct {
	NewEmail string `gorm:"column:email" json:"newEmail"`
	Password string `gorm:"column:password" json:"newPassword"`
}
