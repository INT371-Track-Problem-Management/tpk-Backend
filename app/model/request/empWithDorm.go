package request

type AddEmpInDorm struct {
	DormId     int ` json:"dormId" gorm:"column:dormId"`
	EmployeeId int `json:"employeeId" gorm:"column:employeeId"`
}
