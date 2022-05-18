package entity

type ReviewReports struct {
	ReviewId int    `json:"reviewId" gorm:"column:reviewId"`
	Des      string `json:"des" gorm:"column:des"`
	ReportId int    `json:"reportId" gorm:"column:reportId"`
}
