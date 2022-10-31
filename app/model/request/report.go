package request

type Report struct {
	ReportId int `json:"reportId" gorm:"column:reportId"`
}

type ReportByCreatedBy struct {
	CreateBy int `json:"createBy" gorm:"column:createBy"`
}

type ReportInsert struct {
	Title            string `json:"title" gorm:"column:title"`
	CategoriesReport string `json:"categoriesReport" gorm:"column:categoriesReport"`
	ReportDes        string `json:"reportDes" gorm:"column:reportDes"`
	Status           string `json:"status" gorm:"column:status"`
	UpdateAt         string `json:"updateAt" gorm:"column:updateAt"`
	UpdateBy         int    `json:"updateBy" gorm:"column:updateBy"`
	CreateAt         string `json:"createdAt" gorm:"column:createAt"`
	CreateBy         int    `json:"createdBy" gorm:"column:createBy"`
	RoomId           int    `json:"roomId" gorm:"column:roomId"`
}

type ReportChangeStatus struct {
	ReportId int    `json:"reportId" gorm:"column:reportId"`
	Status   string `json:"status" gorm:"column:status"`
	UpdateBy int    `json:"updateBy" gorm:"column:updateBy"`
}

type EndReport struct {
	ReportId    int    `json:"reportId" gorm:"column:reportId"`
	Description string `json:"description" gorm:"column:des"`
	Score       int    `json:"score" gorm:"column:score"`
	CreatedBy   int    `json:"createdBy" gorm:"column:createdBy"`
}
