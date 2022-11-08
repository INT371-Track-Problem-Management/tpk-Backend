package request

type AssignReport struct {
	EngageId     int `json:"engageId" gorm:"column:engageId"`
	ReportId     int `json:"reportId" gorm:"column:reportId"`
	MaintainerId int `json:"maintainerId" gorm:"column:maintainerId"`
	UpdateBy     int `json:"updateBy" gorm:"column:updateBy"`
}
