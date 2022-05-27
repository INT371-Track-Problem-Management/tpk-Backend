package entity

type HistoryCus struct {
	HistoryId   int    `json:"historyId" gorm:"column:historyId"`
	DateOfEntry string `json:"dateOfEntry" gorm:"column:dateOfEntry"`
	DateOfIssue string `json:"dateOfIssue" gorm:"column:dateOfIssue"`
	RoomId      int    `json:"roomId" gorm:"column:roomId"`
	CustomerId  int    `json:"customerId" gorm:"column:customerId"`
}
