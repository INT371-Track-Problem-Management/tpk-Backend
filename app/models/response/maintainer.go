package response

import "tpk-backend/app/models/model"

type Maintainer struct {
	MaintainerId int    `json:"maintainerId" gorm:"column:maintainerId"`
	Fname        string `json:"fname" gorm:"column:fname"`
	Phone        string `json:"phone" gorm:"column:phone"`
}

type OverviewMaintainer struct {
	MaintainerId int                        `json:"maintainerId" gorm:"column:maintainerId"`
	Fname        string                     `json:"fname" gorm:"column:fname"`
	Lname        string                     `json:"lname" gorm:"column:lname"`
	Phone        string                     `json:"phone" gorm:"column:phone"`
	Overview     []model.OverviewMaintainer `json:"overview" gorm:"column:overview"`
}
