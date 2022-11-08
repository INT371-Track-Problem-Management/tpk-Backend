package request

type ReportStatus struct {
	ReportId  int    `json:"reportId" gorm:"column:reportId"`
	Status    string `json:"status" gorm:"column:status"`
	Detail    string `json:"detail" gorm:"column:detail"`
	UpdateAt  string `json:"updateAt" gorm:"column:updateAt"`
	UpdateBy  int    `json:"updateBy" gorm:"column:updateBy"`
	CreatedAt string `json:"createdAt" gorm:"column:createAt"`
}
