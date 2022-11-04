package model

type ReportEngage struct {
	EngageId     int    `json:"engageId" gorm:"column:engageId"`
	Step         int    `json:"step" gorm:"column:step"`
	SelectedDate int    `json:"selectedDate" gorm:"column:selectedDate"`
	ReportId     int    `json:"reportId" gorm:"column:reportId"`
	BuildingId   int    `json:"buildingId" gorm:"column:buildingId"`
	CreateBy     int    `json:"createBy" gorm:"column:createBy"`
	CreateAt     string `json:"createdAt" gorm:"column:createAt"`
	UpdateAt     string `json:"updateAt" gorm:"column:updateAt"`
	UpdateBy     int    `json:"updateBy" gorm:"column:updateBy"`
	MaintainerId int    `json:"maintainerId" gorm:"column:maintainerId"`
}

type InsertReportEngage struct {
	Step       int    `json:"step" gorm:"column:step"`
	ReportId   int    `json:"reportId" gorm:"column:reportId"`
	BuildingId int    `json:"buildingId" gorm:"column:buildingId"`
	CreateBy   int    `json:"createBy" gorm:"column:createBy"`
	CreateAt   string `json:"createdAt" gorm:"column:createAt"`
	UpdateAt   string `json:"updateAt" gorm:"column:updateAt"`
	UpdateBy   int    `json:"updateBy" gorm:"column:updateBy"`
}

type SelectedPlanFixDate struct {
	EngageId     int    `json:"engageId" gorm:"column:engageId"`
	SelectedDate int    `json:"selectedDate" gorm:"column:selectedDate"`
	UpdateBy     int    `json:"updateBy" gorm:"column:updateBy"`
	UpdateAt     string `json:"updateAt" gorm:"column:updateAt"`
}
