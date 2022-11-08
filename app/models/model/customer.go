package model

type Customer struct {
	CustomerId  int    `json:"customerId" gorm:"column:customerId"`
	Email       string `json:"email" gorm:"column:email"`
	Fname       string `json:"fname" gorm:"column:fname"`
	Lname       string `json:"lname" gorm:"column:lname"`
	Sex         string `json:"sex" gorm:"column:sex"`
	DateOfBirth string `json:"dateOfBirth" gorm:"column:dateOfBirth"`
	Age         int    `json:"age" gorm:"column:age"`
	Phone       string `json:"phone" gorm:"column:phone"`
	Address     string `json:"address" gorm:"column:address"`
	Status      string `json:"status" gorm:"column:status"`
	CreateAt    string `json:"createAt" gorm:"column:createAt"`
	UpdateAt    string `json:"updateAt" gorm:"column:updateAt"`
	UpdateBy    int    `json:"updateBy" gorm:"column:updateBy"`
}

type CustomerEditProfile struct {
	Fname       string `json:"fname" gorm:"column:fname"`
	Lname       string `json:"lname" gorm:"column:lname"`
	Sex         string `json:"sex" gorm:"column:sex"`
	DateOfBirth string `json:"dateOfBirth" gorm:"column:dateOfBirth"`
	Age         int    `json:"age" gorm:"column:age"`
	Phone       string `json:"phone" gorm:"column:phone"`
	Address     string `json:"address" gorm:"column:address"`
	UpdateAt    string `json:"updateAt" gorm:"column:updateAt"`
	UpdateBy    int    `json:"updateBy" gorm:"column:updateBy"`
}
