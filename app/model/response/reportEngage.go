package response

import "tpk-backend/app/model/entity"

type ReportEngage struct {
	EngageId     int    `json:"engageId" gorm:"column:engageId"`
	Date1        string `json:"date1" gorm:"column:date1"`
	Date2        string `json:"date2" gorm:"column:date2"`
	Date3        string `json:"date3" gorm:"column:date3"`
	Date4        string `json:"date4" gorm:"column:date4"`
	SelectedDate string `json:"selectedDate" gorm:"column:selectedDate"`
	ReportId     int    `json:"reportId" gorm:"column:reportId"`
	DormId       int    `json:"dormId" gorm:"column:dormId"`
	UpdatedBy    int    `json:"updatedBy" gorm:"column:updatedBy"`
	SelectedBy   string `json:"selectedBy" gorm:"column:selectedBy"`
}

type ReportEngageAll struct {
	Data []entity.ReportEngage
}

type ReportEngageJoinReport struct {
	EngageId         int    `json:"engageId" gorm:"column:engageId"`
	Date1            string `json:"date1" gorm:"column:date1"`
	Date2            string `json:"date2" gorm:"column:date2"`
	Date3            string `json:"date3" gorm:"column:date3"`
	Date4            string `json:"date4" gorm:"column:date4"`
	SelectedDate     int    `json:"selectedDate" gorm:"column:selectedDate"`
	ReportId         int    `json:"reportId" gorm:"column:reportId"`
	Title            string `json:"title" gorm:"column:title"`
	CategoriesReport string `json:"categoriesReport" gorm:"column:categoriesReport"`
	ReportDes        string `json:"reportDes" gorm:"column:reportDes"`
	Status           string `json:"status" gorm:"column:status"`
	SuccessDate      string `json:"successDate" gorm:"column:successDate"`
	ReportDate       string `json:"reportDate" gorm:"column:reportDate"`
	CreatedBy        int    `json:"createdBy" gorm:"column:createdBy"`
}
