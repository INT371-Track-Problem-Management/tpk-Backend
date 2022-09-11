package entity

type Employee struct {
	EmployeeId  int    `json:"employeeId" gorm:"column:employeeId"`
	Fname       string `json:"fname" gorm:"column:fname"`
	Lname       string `json:"lname" gorm:"column:lname"`
	Phone       string `json:"phone" gorm:"column:phone"`
	Address     string `json:"address" gorm:"column:address"`
	Sex         string `json:"sex" gorm:"column:sex"`
	Email       string `json:"email" gorm:"column:email"`
	Age         int    `json:"age" gorm:"column:age"`
	DateOfBirth string `json:"dateOfBirth" gorm:"column:dateOfBirth"`
	Position    string `json:"position" gorm:"column:position"`
}

type EmployeeRegis struct {
	Fname       string `json:"fname" gorm:"column:fname"`
	Lname       string `json:"lname" gorm:"column:lname"`
	Phone       string `json:"phone" gorm:"column:phone"`
	Address     string `json:"address" gorm:"column:address"`
	Sex         string `json:"sex" gorm:"column:sex"`
	Email       string `json:"email" gorm:"column:email"`
	Age         int    `json:"age" gorm:"column:age"`
	DateOfBirth string `json:"dateOfBirth" gorm:"column:dateOfBirth"`
	Position    string `json:"position" gorm:"column:position"`
}
