package entity

type ReportEngage struct {
	EngageId     int    `json:"engageId" gorm:"column:engageId"`
	Date1        string `json:"date1" gorm:"column:date1"`
	Date2        string `json:"date2" gorm:"column:date2"`
	Date3        string `json:"date3" gorm:"column:date3"`
	Date4        string `json:"date4" gorm:"column:date4"`
	SelectedDate string `json:"selectedDate" gorm:"column:selectedDate"`
	ReportId     int    `json:"reportId" gorm:"column:reportId"`
	BuildingId   int    `json:"buildingId" gorm:"column:buildingId"`
	CreateBy     int    `json:"createBy" gorm:"column:createBy"`
	CreateAt     string `json:"createdAt" gorm:"column:createAt"`
	UpdateAt     string `json:"updateAt" gorm:"column:updateAt"`
	UpdateBy     int    `json:"updateBy" gorm:"column:updateBy"`
	MaintainerId	 int 	`json:"maintainerId" gorm:"column:maintainerId"`
}

type InsertReportEngage struct {
	Date1      string `json:"date1" gorm:"column:date1"`
	Date2      string `json:"date2" gorm:"column:date2"`
	Date3      string `json:"date3" gorm:"column:date3"`
	Date4      string `json:"date4" gorm:"column:date4"`
	ReportId   int    `json:"reportId" gorm:"column:reportId"`
	BuildingId int    `json:"buildingId" gorm:"column:buildingId"`
	CreateBy   int    `json:"createBy" gorm:"column:createBy"`
	CreateAt   string `json:"createdAt" gorm:"column:createAt"`
	UpdateAt   string `json:"updateAt" gorm:"column:updateAt"`
	UpdateBy   int    `json:"updateBy" gorm:"column:updateBy"`
}

type SelectedPlanFixDate struct {
	EngageId     int    `json:"engageId" gorm:"column:engageId"`
	SelectedDate string `json:"selectedDate" gorm:"column:selectedDate"`
	UpdateBy     int    `json:"updateBy" gorm:"column:updateBy"`
	UpdateAt     string `json:"updateAt" gorm:"column:updateAt"`
}
