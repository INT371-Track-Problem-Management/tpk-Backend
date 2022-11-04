package model

type SaveToken struct {
	Token  string `json:"token" gorm:"column:token"`
	Status string `json:"status" gorm:"column:status"`
}
