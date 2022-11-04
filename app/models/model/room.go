package model

type Room struct {
	RoomId      int    `json:"roomId" gorm:"column:roomId"`
	RoomNum     string `json:"roomNum" gorm:"column:roomNum"`
	Floors      int    `json:"floors" gorm:"column:floors"`
	Description string `json:"description" gorm:"column:description"`
	BuildingId  int    `json:"buildingId" gorm:"column:buildingId"`
	Status      string `json:"status" gorm:"column:status"`
	UpdateAt    string `json:"updateAt" gorm:"column:updateAt"`
	UpdateBy    int    `json:"updateBy" gorm:"column:updateBy"`
	CreateAt    string `json:"createAt" gorm:"column:createAt"`
}

type RoomInsert struct {
	RoomNum     string `json:"roomNum" gorm:"column:roomNum"`
	Floors      int    `json:"floors" gorm:"column:floors"`
	Description string `json:"description" gorm:"column:description"`
	BuildingId  int    `json:"buildingId" gorm:"column:buildingId"`
	Status      string `json:"status" gorm:"column:status"`
	UpdateAt    string `json:"updateAt" gorm:"column:updateAt"`
	UpdateBy    int    `json:"updateBy" gorm:"column:updateBy"`
	CreateAt    string `json:"createAt" gorm:"column:createAt"`
}

type RoomsStatus struct {
	RoomId   int    `json:"roomId" gorm:"column:roomId"`
	Status   string `json:"status" gorm:"column:status"`
	UpdateAt string `json:"updateAt" gorm:"column:updateAt"`
	UpdateBy int    `json:"updateBy" gorm:"column:updateBy"`
}

type RoomByFloors struct {
	RoomId  int    `json:"roomId" gorm:"column:roomId"`
	RoomNum string `json:"roomNum" gorm:"column:roomNum"`
	Status  string `json:"status" gorm:"column:status"`
}
