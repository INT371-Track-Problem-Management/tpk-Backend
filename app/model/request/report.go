package request

type Report struct {
	ReportId int `json:"reportId" gorm:"column:reportId"`
}

type ReportInsert struct {
	ReportId         int    `json:"reportId" gorm:"column:reportId"`
	Title            string `json:"title" gorm:"column:title"`
	CategoriesReport string `json:"categoriesReport" gorm:"column:categoriesReport"`
	ReportDes        string `json:"reportDes" gorm:"column:reportDes"`
	Status           string `json:"status" gorm:"column:status"`
	ReviewId         int    `json:"reviewId" gorm:"column:reviewId"`
}

type ReportChangeStatus struct {
	ReportId int    `json:"reportId" gorm:"column:reportId"`
	Status   string `json:"status" gorm:"column:status"`
}
