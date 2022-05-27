package entity

type Report struct {
	ReportId         int    `json:"reportId" gorm:"column:reportId"`
	Title            string `json:"title" gorm:"column:title"`
	CategoriesReport string `json:"categoriesReport" gorm:"column:categoriesReport"`
	ReportDes        string `json:"reportDes" gorm:"column:reportDes"`
	Status           string `json:"status" gorm:"column:status"`
	SuccessDate      string `json:"successDate" gorm:"column:successDate"`
	ReportDate       string `json:"reportDate" gorm:"column:reportDate"`
	CreatedBy        int    `json:"createdBy" gorm:"column:createdBy"`
}

type ReportInsert struct {
	Title            string `json:"title" gorm:"column:title"`
	CategoriesReport string `json:"categoriesReport" gorm:"column:categoriesReport"`
	ReportDes        string `json:"reportDes" gorm:"column:reportDes"`
	Status           string `json:"status" gorm:"column:status"`
	ReportDate       string `json:"reportDate" gorm:"column:reportDate"`
	CreatedBy        int    `json:"createdBy" gorm:"column:createdBy"`
}
