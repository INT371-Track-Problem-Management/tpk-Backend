package entity

type RoomWithCustomer struct {
	Id         int    `json:"id" gorm:"column:Id"`
	RoomId     int    `json:"roomId" gorm:"column:roomId"`
	CustomerId int    `json:"customerId" gorm:"column:customerId"`
	Status     string `json:"status" gorm:"column:status"`
}
