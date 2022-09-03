package request

type Report struct {
	ReportId int `json:"reportId" gorm:"column:reportId"`
}

type ReportByCreatedBy struct {
	CreatedBy int `json:"createdBy" gorm:"column:createdBy"`
}

type ReportInsert struct {
	Title            string `json:"title" gorm:"column:title"`
	CategoriesReport string `json:"categoriesReport" gorm:"column:categoriesReport"`
	ReportDes        string `json:"reportDes" gorm:"column:reportDes"`
	Status           string `json:"status" gorm:"column:status"`
	CreatedBy        int    `json:"createdBy" gorm:"column:createdBy"`
}

type ReportChangeStatus struct {
	ReportId   int    `json:"reportId" gorm:"column:reportId"`
	Status     string `json:"status" gorm:"column:status"`
	EmployeeId int    `json:"employeeId" gorm:"column:employeeId"`
}

type EndReport struct {
	ReportId    int    `json:"reportId" gorm:"column:reportId"`
	Description string `json:"description" gorm:"column:des"`
	Score       int    `json:"score" gorm:"column:score"`
	CreatedBy   int    `json:"createdBy" gorm:"column:createdBy"`
}
