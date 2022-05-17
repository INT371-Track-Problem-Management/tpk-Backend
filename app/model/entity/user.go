package entity

type User struct {
	Username string `gorm:"column:username" json:"username"`
	Password string `gorm:"column:password" json:"password"`
	Role     string `gorm:"column:role" json:"role"`
}
