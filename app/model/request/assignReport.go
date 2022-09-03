package request

type AssignReport struct {
	AssignDate   string `json:"assignDate" gorm:"column:assignDate"`
	MaintainerId int    `json:"maintainerId" gorm:"column:maintainerId"`
	EmployeeId   int    `json:"employeeId" gorm:"column:employeeId"`
	ReportId     int    `json:"reportId" gorm:"column:reportId"`
}
