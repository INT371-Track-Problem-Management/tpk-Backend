package request

type RoomAddCustomer struct {
	RoomId     int `json:"roomId" gorm:"column:roomId"`
	CustomerId int `json:"customerId" gorm:"column:customerId"`
	BuildingId int `json:"buildingId" gorm:"column:buildingId"`
	UpdateBy   int `json:"updateBy" gorm:"column:updateBy"`
}
