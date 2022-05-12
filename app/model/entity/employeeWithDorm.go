package entity

type EmployeeWithDorm struct {
	EmployeeId int `json:"employeeId" gorm:"column:employeeId"`
	DormId     int `json:"dormId" gorm:"column:dormId"`
}
