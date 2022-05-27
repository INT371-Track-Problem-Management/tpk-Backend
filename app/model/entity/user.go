package entity

type User struct {
	Email    string `gorm:"column:email" json:"email"`
	Password string `gorm:"column:password" json:"password"`
	Role     string `gorm:"column:role" json:"role"`
}
