package response

import "tpk-backend/app/models/model"

type ReportList struct {
	ReportId         int    `json:"reportId" gorm:"column:reportId"`
	Title            string `json:"title" gorm:"column:title"`
	CategoriesReport string `json:"categoriesReport" gorm:"column:categoriesReport"`
	ReportDes        string `json:"reportDes" gorm:"column:reportDes"`
	Status           string `json:"status" gorm:"column:status"`
	RoomId           int    `json:"roomId" gorm:"column:roomId"`
	RoomNum          string `json:"roomNum" gorm:"column:roomNum"`
	UpdateAt         string `json:"updateAt" gorm:"column:updateAt"`
	UpdateBy         int    `json:"updateBy" gorm:"column:updateBy"`
	CreateAt         string `json:"createdAt" gorm:"column:createAt"`
	CreateBy         int    `json:"createBy" gorm:"column:createBy"`
	ImageId          string `json:"imageId" gorm:"column:imageId"`

	EngageId     int                `json:"-" gorm:"column:engageId"`
	SelectedDate int                `json:"-" gorm:"column:selectedDate"`
	Engage       model.ReportEngage `json:"-" gorm:"foreignKey:ReportId;references:EngageId"`
	Id           int                `json:"-" gorm:"column:id"`
	Step         string             `json:"step" gorm:"column:step","foreignKey:Id;references:SelectedDate"`
	FixDate      string             `json:"fixDate" gorm:"column:date","foreignKey:Id;references:SelectedDate"`
}

type ReportDetailById struct {
	ReportId         int    `json:"reportId" gorm:"column:reportId"`
	Title            string `json:"title" gorm:"column:title"`
	CategoriesReport string `json:"categoriesReport" gorm:"column:categoriesReport"`
	ReportDes        string `json:"reportDes" gorm:"column:reportDes"`
	Status           string `json:"status" gorm:"column:status"`
	RoomId           int    `json:"roomId" gorm:"column:roomId"`
	RoomNum          string `json:"roomNum" gorm:"column:roomNum"`
	UpdateAt         string `json:"updateAt" gorm:"column:updateAt"`
	UpdateBy         int    `json:"updateBy" gorm:"column:updateBy"`
	CreateAt         string `json:"createdAt" gorm:"column:createAt"`
	CreateBy         int    `json:"createBy" gorm:"column:createBy"`
	ImageId          string `json:"imageId" gorm:"column:imageId"`
}

type Year struct {
	Year []int `json:"year" gorm:"column:year"`
}
