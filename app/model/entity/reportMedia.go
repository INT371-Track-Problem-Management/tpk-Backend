package entity

type ReportMedia struct {
	Id        int    `json:"id" gorm:"column:id"`
	Url       string `json:"url" gorm:"column:url"`
	ReportId  string `json:"reportId" gorm:"column:reportId"`
	ConfigKey string `json:"configKey" gorm:"column:configKey"`
	CreateAt  string `json:"createAt" gorm:"column:createAt"`
	UpdateAt  string `json:"updateAt" gorm:"column:updateAt"`
	UpdateBy  int    `json:"updateBy" gorm:"column:updateBy"`
}
