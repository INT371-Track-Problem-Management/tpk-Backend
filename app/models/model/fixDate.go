package model

type Fixdate struct {
	Id       int    `json:"id" gorm:"column:id"`
	Date     string `json:"date" gorm:"column:date"`
	Step     int    `json:"step" gorm:"column:step"`
	CreateAt string `json:"createAt" gorm:"column:createAt"`
	EngageId int    `json:"engageId" gorm:"column:engageId"`
}

type CreateFixdate struct {
	Date     string `json:"Date" gorm:"column:Date"`
	Step     int    `json:"step" gorm:"column:step"`
	CreateAt string `json:"createAt" gorm:"column:createAt"`
	EngageId int    `json:"engageId" gorm:"column:engageId"`
}
