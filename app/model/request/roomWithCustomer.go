package request

type RoomAddCustomer struct {
	RoomId     int `json:"roomId" gorm:"column:roomId"`
	CustomerId int `json:"customerId" gorm:"column:customerId"`
}
