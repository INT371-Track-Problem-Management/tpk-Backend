package request

type ReportEngage struct {
	EngageId int `json:"engageId" gorm:"column:engageId"`
	Step     int `json:"step" gorm:"column:step"`
	Dates    []struct {
		Date string `json:"date" gorm:"column:date"`
	} `json:"dates" gorm:"column:dates"`
	UpdatedBy int `json:"updateBy" gorm:"column:updateBy"`
}

type SelectedPlanFixDate struct {
	EngageId     int `json:"engageId" gorm:"column:engageId"`
	SelectedDate int `json:"selectedDate" gorm:"column:selectedDate"`
	UpdateBy     int `json:"updateBy" gorm:"column:updateBy"`
}
