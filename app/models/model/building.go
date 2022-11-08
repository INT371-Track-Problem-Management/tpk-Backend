package model

type Building struct {
	BuildingId   int    `json:"buildingId" gorm:"column:buildingId"`
	BuildingName string `json:"buildingName" gorm:"column:buildingName"`
	Address      string `json:"address" gorm:"column:address"`
	Phone        string `json:"phone" gorm:"column:phone"`
	Email        string `json:"email" gorm:"column:email"`
	CreateAt     string `json:"createdAt" gorm:"column:createAt"`
	UpdateAt     string `json:"updateAt" gorm:"column:updateAt"`
	UpdateBy     int    `json:"createBy" gorm:"column:createBy"`
}

type BuildingInsert struct {
	BuildingName string `json:"buildingName" gorm:"column:buildingName"`
	CreateAt     string `json:"createdAt" gorm:"column:createAt"`
	UpdateAt     string `json:"updateAt" gorm:"column:updateAt"`
	UpdateBy     int    `json:"createBy" gorm:"column:updateBy"`
}
