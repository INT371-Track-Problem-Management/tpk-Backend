package request

type Report struct {
	ReportId int `json:"reportId" gorm:"column:reportId"`
}
