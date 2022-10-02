package entity

type RoomWithCustomer struct {
	Id         int    `json:"id" gorm:"column:id"`
	RoomId     int    `json:"roomId" gorm:"column:roomId"`
	CustomerId int    `json:"customerId" gorm:"column:customerId"`
	Status     string `json:"status" gorm:"column:status"`
	DormId     int    `json:"dormId" gorm:"column:dormId"`
}

type RoomJoinDorm struct {
	Id          int    `json:"id" gorm:"column:id"`
	RoomId      int    `json:"roomId" gorm:"column:roomId"`
	CustomerId  int    `json:"customerId" gorm:"column:customerId"`
	Status      string `json:"status" gorm:"column:status"`
	UpdateAt    string `json:"updateAt" gorm:"column:updateAt"`
	UpdateBy    int    `json:"updateBy" gorm:"column:updateBy"`
	CreateAt    string `json:"createAt" gorm:"column:createAt"`
	RoomNum     string `json:"roomNum" gorm:"column:roomNum"`
	Floors      int    `json:"floors" gorm:"column:floors"`
	Description string `json:"description" gorm:"column:description"`
	BuildingId  int    `json:"buildingId" gorm:"column:buildingId"`
}
