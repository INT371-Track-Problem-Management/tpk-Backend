package request

type ReportEngage struct {
	Date1      string `json:"date1" gorm:"column:date1"`
	Date2      string `json:"date2" gorm:"column:date2"`
	Date3      string `json:"date3" gorm:"column:date3"`
	Date4      string `json:"date4" gorm:"column:date4"`
	ReportId   int    `json:"reportId" gorm:"column:reportId"`
	DormId     int    `json:"dormId" gorm:"column:dormId"`
	UpdatedBy  int    `json:"updatedBy" gorm:"column:updatedBy"`
	SelectedBy int    `json:"selectedBy" gorm:"column:selectedBy"`
}

type ReportEngageById struct {
	EngageId int `json:"engageId" gorm:"column:engageId"`
}

type SelectedPlanFixDate struct {
	EngageId     int    `json:"engageId" gorm:"column:engageId"`
	SelectedDate string `json:"selectedDate" gorm:"column:selectedDate"`
}
