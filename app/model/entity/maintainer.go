package entity

type Maintainer struct {
	MaintainerId string `json:"maintainerId" gorm:"column:maintainerId"`
	Fname        string `json:"fname" gorm:"column:fname"`
	Phone        string `json:"phone" gorm:"column:phone"`
}
