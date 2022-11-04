package response

import "tpk-backend/app/models/model"

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
	MaintainerId int    `json:"maintainerId" gorm:"column:maintainerId"`
}

type ReportEngageAll struct {
	Data []model.ReportEngage
}

type ReportEngageJoinReport struct {
	EngageId         int    `json:"engageId" gorm:"column:engageId"`
	Date1            string `json:"date1" gorm:"column:date1"`
	Date2            string `json:"date2" gorm:"column:date2"`
	Date3            string `json:"date3" gorm:"column:date3"`
	Date4            string `json:"date4" gorm:"column:date4"`
	SelectedDate     string `json:"selectedDate" gorm:"column:selectedDate"`
	ReportId         int    `json:"reportId" gorm:"column:reportId"`
	Title            string `json:"title" gorm:"column:title"`
	CategoriesReport string `json:"categoriesReport" gorm:"column:categoriesReport"`
	ReportDes        string `json:"reportDes" gorm:"column:reportDes"`
	Status           string `json:"status" gorm:"column:status"`
	UpdateAt         string `json:"updateAt" gorm:"column:updateAt"`
	UpdateBy         int    `json:"updateBy" gorm:"column:updateBy"`
	CreateAt         string `json:"createdAt" gorm:"column:createAt"`
	CreateBy         int    `json:"createdBy" gorm:"column:createBy"`
	RoomNum          string `json:"roomNum" gorm:"column:roomNum"`
	BuildingId       int    `json:"buildingId" gorm:"column:buildingId"`
	MaintainerId     int    `json:"maintainerId" gorm:"column:maintainerId"`
}
