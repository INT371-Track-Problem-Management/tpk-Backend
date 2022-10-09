package response

type Report struct {
	ReportId         int    `json:"reportId" gorm:"column:reportId"`
	Title            string `json:"title" gorm:"column:title"`
	CategoriesReport string `json:"categoriesReport" gorm:"column:categoriesReport"`
	ReportDes        string `json:"reportDes" gorm:"column:reportDes"`
	Status           string `json:"status" gorm:"column:status"`
	UpdateAt         string `json:"updateAt" gorm:"column:updateAt"`
	UpdateBy         int    `json:"updateBy" gorm:"column:updateBy"`
	CreateAt         string `json:"createdAt" gorm:"column:createAt"`
	CreateBy         int    `json:"createdBy" gorm:"column:createBy"`
}

type Year struct {
	Year []int `json:"year" gorm:"column:year"`
}
