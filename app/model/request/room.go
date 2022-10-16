package request

type RoomsStatus struct {
	RoomId   int    `json:"roomId" gorm:"column:roomId"`
	Status   string `json:"status" gorm:"column:status"`
	UpdateBy int    `json:"updateBy" gorm:"column:updateBy"`
}

type RoomInsert struct {
	BuildingId int `json:"buildingId" gorm:"column:buildingId"`
	Rooms      []struct {
		RoomNum     string `json:"roomNum" gorm:"column:roomNum"`
		Floors      int    `json:"floors" gorm:"column:floors"`
		Description string `json:"description" gorm:"column:description"`
	} `json:"rooms"`
	UpdateBy int `json:"updateBy" gorm:"column:updateBy"`
}
