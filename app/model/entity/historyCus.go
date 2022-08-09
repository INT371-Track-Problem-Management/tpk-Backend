package entity

type HistoryReport struct {
	HistoryId   int    `json:"historyId" gorm:"column:historyId"`
	ReportId    int    `json:"reportId" gorm:"column:reportId"`
	ReportDate  string `json:"dateOfEntry" gorm:"column:dateOfEntry"`
	DateOfIssue string `json:"dateOfIssue" gorm:"column:dateOfIssue"`
	RoomId      int    `json:"roomId" gorm:"column:roomId"`
	CustomerId  int    `json:"customerId" gorm:"column:customerId"`
	EmployeeId  int    `json:"employeeId" gorm:"column:employeeId"`
}
