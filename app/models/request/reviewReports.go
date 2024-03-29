package request

type ReviewReports struct {
	Des      string `json:"des" gorm:"column:des"`
	ReportId int    `json:"reportId" gorm:"column:reportId"`
}

type EndJobReport struct {
	Des          string `json:"des" gorm:"column:des"`
	ReportId     int    `json:"reportId" gorm:"column:reportId"`
	Score        int    `json:"score" gorm:"column:score"`
	MaintainerId int    `json:"maintainerId" gorm:"column:maintainerId"`
	UpdateBy     int    `json:"updateBy" gorm:"column:updateBy"`
}
