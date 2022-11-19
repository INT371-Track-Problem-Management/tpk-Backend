package model

type Stat struct {
	TotalReport    int `json:"total_report" gorm:"column:total_report"`
	Electric       int `json:"electric" gorm:"column:electric"`
	Water          int `json:"water" gorm:"column:water"`
	ElectricDevice int `json:"electric_device" gorm:"column:electric_device"`
	WaterMachine   int `json:"water_machine" gorm:"column:water_machine"`
	Furniture      int `json:"furniture" gorm:"column:furniture"`
	Building       int `json:"building" gorm:"column:building"`
	Other          int `json:"other" gorm:"column:other"`
}

type StatMaintainer struct {
	MaintainerId int     `json:"maintainerId" gorm:"column:maintainerId"`
	Fname        string  `json:"fname" gorm:"column:fname"`
	Lname        string  `json:"lname" gorm:"column:lname"`
	Phone        string  `json:"phone" gorm:"column:phone"`
	Total        string  `json:"total" gorm:"column:total"`
	Average      float64 `json:"average" gorm:"column:average"`
}

type OverviewMaintainer struct {
	Score int    `json:"score" gorm:"column:score"`
	Des   string `json:"des" gorm:"column:des"`
}
