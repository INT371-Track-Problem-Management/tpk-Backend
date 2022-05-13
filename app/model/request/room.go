package request

type RoomsStatus struct {
	RoomId int    `json:"roomId" gorm:"column:roomId"`
	Status string `json:"status" gorm:"column:status"`
}
