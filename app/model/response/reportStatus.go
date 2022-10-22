package response

type ReportStatus struct {
	StatusId  int    `json:"statusId" gorm:"column:statusId"`
	ReportId  int    `json:"reportId" gorm:"column:reportId"`
	Status    string `json:"status" gorm:"column:status"`
	CreatedAt string `json:"createdAt" gorm:"column:createAt"`
}
