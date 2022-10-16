package response

type Room struct {
	RoomId      int    `json:"roomId" gorm:"column:roomId"`
	RoomNum     string `json:"roomNum" gorm:"column:roomNum"`
	Floors      int    `json:"floors" gorm:"column:floors"`
	Description string `json:"description" gorm:"column:description"`
	Status      string `json:"status" gorm:"column:status"`
	BuildingId  int    `json:"buildingId" gorm:"column:buildingId"`
}

type RoomByBuildingId struct {
	BuildingId string        `json:"buildingId" gorm:"column:buildingId"`
	Floors     []interface{} `json:"floors" gorm:"column:floors"`
}
