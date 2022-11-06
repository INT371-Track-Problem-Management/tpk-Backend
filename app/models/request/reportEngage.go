package request

type ReportEngage struct {
	Date1      string `json:"date1" gorm:"column:date1"`
	Date2      string `json:"date2" gorm:"column:date2"`
	Date3      string `json:"date3" gorm:"column:date3"`
	Date4      string `json:"date4" gorm:"column:date4"`
	ReportId   int    `json:"reportId" gorm:"column:reportId"`
	BuildingId int    `json:"buildingId" gorm:"column:buildingId"`
	UpdatedBy  int    `json:"updateBy" gorm:"column:updateBy"`
}

type ReportEngageById struct {
	EngageId int `json:"engageId" gorm:"column:engageId"`
}

type ReportEngageByReportId struct {
	ReportId int `json:"reportId" gorm:"column:reportId"`
}

type SelectedPlanFixDate struct {
	EngageId     int    `json:"engageId" gorm:"column:engageId"`
	SelectedDate string `json:"selectedDate" gorm:"column:selectedDate"`
	UpdateBy     int    `json:"updateBy" gorm:"column:updateBy"`
}