package entity

type ReviewReports struct {
	ReviewId int    `json:"reviewId" gorm:"column:reviewId"`
	Des      string `json:"des" gorm:"column:des"`
	ReportId int    `json:"reportId" gorm:"column:reportId"`
}

type EndJobReport struct {
	Des         string `json:"des" gorm:"column:des"`
	ReportId    int    `json:"reportId" gorm:"column:reportId"`
	Score       int    `json:"score" gorm:"column:score"`
	DateOfIssue string `json:"dateOfIssue" gorm:"column:dateOfIssue"`
}
