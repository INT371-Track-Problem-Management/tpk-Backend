package entity

type Maintainer struct {
	MaintainerId int    `json:"maintainerId" gorm:"column:maintainerId"`
	Fname        string `json:"fname" gorm:"column:fname"`
	Lname        string `json:"lname" gorm:"column:lname"`
	Phone        string `json:"phone" gorm:"column:phone"`
	CreatedBy    int    `json:"createdBy" gorm:"column:createdBy"`
	CreatedDate  string `json:"createdDate" gorm:"column:createdDate"`
}

type AddMaintainer struct {
	Fname       string `json:"fname" gorm:"column:fname"`
	Lname       string `json:"lname" gorm:"column:lname"`
	Phone       string `json:"phone" gorm:"column:phone"`
	CreatedBy   int    `json:"createdBy" gorm:"column:createdBy"`
	CreatedDate string `json:"createdDate" gorm:"column:createdDate"`
}
