package entity

type ReportStatus struct {
	StatusId  int    `json:"statusId" gorm:"column:statusId"`
	ReportId  int    `json:"reportId" gorm:"column:reportId"`
	Status    string `json:"status" gorm:"column:status"`
	UpdateAt  string `json:"updateAt" gorm:"column:updateAt"`
	UpdateBy  int    `json:"updateBy" gorm:"column:updateBy"`
	CreatedAt string `json:"createdAt" gorm:"column:createAt"`
}
