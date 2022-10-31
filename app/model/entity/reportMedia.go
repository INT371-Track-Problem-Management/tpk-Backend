package entity

type ReportMedia struct {
	Id       int    `json:"id" gorm:"column:id"`
	ReportId int    `json:"reportId" gorm:"column:reportId"`
	Name     string `json:"name" gorm:"column:name"`
	Url      string `json:"url" gorm:"column:url"`
	CreateAt string `json:"createAt" gorm:"column:createAt"`
	UpdateAt string `json:"updateAt" gorm:"column:updateAt"`
}

type ReportMediaInsert struct {
	ReportId int    `json:"reportId" gorm:"column:reportId"`
	Name     string `json:"name" gorm:"column:name"`
	Url      string `json:"url" gorm:"column:url"`
	CreateAt string `json:"createAt" gorm:"column:createAt"`
	UpdateAt string `json:"updateAt" gorm:"column:updateAt"`
}
