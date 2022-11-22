package response

type Customer struct {
	CustomerId  int    `json:"customerId" gorm:"column:customerId"`
	Email       string `json:"email" gorm:"column:email"`
	Password    string `json:"password" gorm:"column:password"`
	Fname       string `json:"fname" gorm:"column:fname"`
	Lname       string `json:"lname" gorm:"column:lname"`
	Sex         string `json:"sex" gorm:"column:sex"`
	DateOfBirth string `json:"dateOfBirth" gorm:"column:dateOfBirth"`
	Age         int    `json:"age" gorm:"column:age"`
	Phone       string `json:"phone" gorm:"column:phone"`
}

type CustomerProfile struct {
	CustomerId  int    `json:"customerId" gorm:"column:customerId"`
	Email       string `json:"email" gorm:"column:email"`
	Fname       string `json:"fname" gorm:"column:fname"`
	Lname       string `json:"lname" gorm:"column:lname"`
	Sex         string `json:"sex" gorm:"column:sex"`
	DateOfBirth string `json:"dateOfBirth" gorm:"column:dateOfBirth"`
	Age         int    `json:"age" gorm:"column:age"`
	Phone       string `json:"phone" gorm:"column:phone"`
}

type ListCustomer struct {
	CustomerId int    `json:"customerId" gorm:"column:customerId"`
	Email      string `json:"email" gorm:"column:email"`
	Fname      string `json:"fname" gorm:"column:fname"`
	Lname      string `json:"lname" gorm:"column:lname"`
	RoomId     int    `json:"roomId" gorm:"column:roomId"`
	RoomNum    string `json:"roomNum" gorm:"column:roomNum"`
	Floors     int    `json:"floors" gorm:"column:floors"`
	BuildingId int    `json:"buildingId" gorm:"column:buildingId"`
}
