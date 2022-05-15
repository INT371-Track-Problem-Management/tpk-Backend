package request

type RoomsStatus struct {
	RoomId int    `json:"roomId" gorm:"column:roomId"`
	Status string `json:"status" gorm:"column:status"`
}

type RoomInsert struct {
	RoomId      int    `json:"roomId" gorm:"column:roomId"`
	RoomNum     string `json:"roomNum" gorm:"column:roomNum"`
	Floors      int    `json:"floors" gorm:"column:floors"`
	Description string `json:"description" gorm:"column:description"`
	DormId      int    `json:"dormId" gorm:"column:dormId"`
}
