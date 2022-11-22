package model

type Maintainer struct {
	MaintainerId int    `json:"maintainerId" gorm:"column:maintainerId"`
	Fname        string `json:"fname" gorm:"column:fname"`
	Lname        string `json:"lname" gorm:"column:lname"`
	Phone        string `json:"phone" gorm:"column:phone"`
	CreateAt     string `json:"createAt" gorm:"column:createAt"`
	UpdateAt     string `json:"updateAt" gorm:"column:updateAt"`
	UpdateBy     int    `json:"updateBy" gorm:"column:updateBy"`
}

type AddMaintainer struct {
	Fname    string `json:"fname" gorm:"column:fname"`
	Lname    string `json:"lname" gorm:"column:lname"`
	Phone    string `json:"phone" gorm:"column:phone"`
	CreateAt string `json:"createAt" gorm:"column:createAt"`
	UpdateAt string `json:"updateAt" gorm:"column:updateAt"`
	UpdateBy int    `json:"updateBy" gorm:"column:updateBy"`
}
