package entity

type RoomWithCustomer struct {
	Id         int    `json:"id" gorm:"column:id"`
	RoomId     int    `json:"roomId" gorm:"column:roomId"`
	CustomerId int    `json:"customerId" gorm:"column:customerId"`
	Status     string `json:"status" gorm:"column:status"`
	BuildingId int    `json:"buildingId" gorm:"column:buildingId"`
}

type RoomJoinBulding struct {
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

type RoomAddCustomer struct {
	RoomId     int    `json:"roomId" gorm:"column:roomId"`
	CustomerId int    `json:"customerId" gorm:"column:customerId"`
	BuildingId int    `json:"buildingId" gorm:"column:buildingId"`
	Status     string `json:"status" gorm:"column:status"`
	CreateAt   string `json:"createAt" gorm:"column:createAt"`
	UpdateAt   string `json:"updateAt" gorm:"column:updateAt"`
	UpdateBy   int    `json:"updateBy" gorm:"column:updateBy"`
}

type RoomRemoveCustomer struct {
	Id       int    `json:"id" gorm:"column:id"`
	Status   string `json:"status" gorm:"column:status"`
	UpdateAt string `json:"updateAt" gorm:"column:updateAt"`
	UpdateBy int    `json:"updateBy" gorm:"column:updateBy"`
}

type RoomWithCustomerId struct {
	RoomId     int    `json:"roomId" gorm:"column:roomId"`
	RoomNum    string `json:"roomNum" gorm:"column:roomNum"`
	BuildingId int    `json:"buildingId" gorm:"column:buildingId"`
	Floors     int    `json:"floors" gorm:"column:floors"`
	Status     string `json:"status" gorm:"column:status"`
}
