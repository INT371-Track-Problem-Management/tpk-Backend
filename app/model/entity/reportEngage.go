package entity

type ReportEngage struct {
	EngageId   int    `json:"engageId" gorm:"column:engageId"`
	Date1      string `json:"date1" gorm:"column:date1"`
	Date2      string `json:"date2" gorm:"column:date2"`
	Date3      string `json:"date3" gorm:"column:date3"`
	Date4      string `json:"date4" gorm:"column:date4"`
	SelectDate int    `json:"selectDate" gorm:"column:selectDate"`
	ReportId   int    `json:"reportId" gorm:"column:reportId"`
}
