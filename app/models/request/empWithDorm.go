package request

type AddEmpInBuilding struct {
	BuildingId int `json:"buildingId" gorm:"column:buildingId"`
	EmployeeId int `json:"employeeId" gorm:"column:employeeId"`
}
