package request

type Report struct {
	ReportId int `json:"reportId" gorm:"column:reportId"`
}

type FillterReport struct {
	RoomId     string
	CustomerId string
}

type ReportByCreatedBy struct {
	CreateBy int `json:"createBy" gorm:"column:createBy"`
}

type ReportInsert struct {
	Title            string `json:"title" gorm:"column:title"`
	CategoriesReport string `json:"categoriesReport" gorm:"column:categoriesReport"`
	ReportDes        string `json:"reportDes" gorm:"column:reportDes"`
	Status           string `json:"status" gorm:"column:status"`
	RoomId           int    `json:"roomId" gorm:"column:roomId"`
	BuildingId       int    `json:"buildingId" gorm:"column:buildingId"`
	Step             int    `json:"step" gorm:"column:step"`
	Dates            []struct {
		Date string `json:"date" gorm:"column:date"`
	} `json:"dates" gorm:"column:dates"`
	UpdateBy int `json:"updateBy" gorm:"column:updateBy"`
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
