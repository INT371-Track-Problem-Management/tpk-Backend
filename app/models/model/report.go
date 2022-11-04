package model

type Report struct {
	ReportId         int    `json:"reportId" gorm:"column:reportId"`
	Title            string `json:"title" gorm:"column:title"`
	CategoriesReport string `json:"categoriesReport" gorm:"column:categoriesReport"`
	ReportDes        string `json:"reportDes" gorm:"column:reportDes"`
	Status           string `json:"status" gorm:"column:status"`
	RoomId           int    `json:"roomId" gorm:"column:roomId"`
	UpdateAt         string `json:"updateAt" gorm:"column:updateAt"`
	UpdateBy         int    `json:"updateBy" gorm:"column:updateBy"`
	CreateAt         string `json:"createdAt" gorm:"column:createAt"`
	CreateBy         int    `json:"createBy" gorm:"column:createBy"`
	RoomNum          string `json:"roomNum" gorm:"column:roomNum"`
	BuildingId       int    `json:"buildingId" gorm:"column:buildingId"`
	SelectedDate     string `json:"selectedDate" gorm:"column:selectedDate"`
}

type ReportInsert struct {
	Title            string `json:"title" gorm:"column:title"`
	CategoriesReport string `json:"categoriesReport" gorm:"column:categoriesReport"`
	ReportDes        string `json:"reportDes" gorm:"column:reportDes"`
	Status           string `json:"status" gorm:"column:status"`
	RoomId           int    `json:"roomId" gorm:"column:roomId"`
	UpdateAt         string `json:"updateAt" gorm:"column:updateAt"`
	UpdateBy         int    `json:"updateBy" gorm:"column:updateBy"`
	CreateAt         string `json:"createdAt" gorm:"column:createAt"`
	CreateBy         int    `json:"createdBy" gorm:"column:createBy"`
}

type EndReport struct {
	ReportId    int    `json:"reportId" gorm:"column:reportId"`
	Description string `json:"description" gorm:"column:des"`
	Score       int    `json:"score" gorm:"column:score"`
	CreatedBy   int    `json:"createdBy" gorm:"column:createdBy"`
	DateOfIssue string `json:"dateOfIssue" gorm:"column:dateOfIssue"`
}

type ReportChangeStatus struct {
	ReportId int    `json:"reportId" gorm:"column:reportId"`
	Status   string `json:"status" gorm:"column:status"`
	UpdateAt string `json:"updateAt" gorm:"column:updateAt"`
	UpdateBy int    `json:"updateBy" gorm:"column:updateBy"`
}

type ReportByRoomId struct {
	ReportId         int    `json:"reportId" gorm:"column:reportId"`
	Title            string `json:"title" gorm:"column:title"`
	CategoriesReport string `json:"categoriesReport" gorm:"column:categoriesReport"`
	ReportDes        string `json:"reportDes" gorm:"column:reportDes"`
	Status           string `json:"status" gorm:"column:status"`
	UpdateAt         string `json:"updateAt" gorm:"column:updateAt"`
	UpdateBy         int    `json:"updateBy" gorm:"column:updateBy"`
	CreateAt         string `json:"createdAt" gorm:"column:createAt"`
	CreateBy         int    `json:"createBy" gorm:"column:createBy"`
	RoomId           int    `json:"roomId" gorm:"column:roomId"`
	SelectedDate     string `json:"selectedDate" gorm:"column:selectedDate"`
}

type ReportJoinEngage struct {
	ReportId         int    `json:"reportId" gorm:"column:reportId"`
	Title            string `json:"title" gorm:"column:title"`
	CategoriesReport string `json:"categoriesReport" gorm:"column:categoriesReport"`
	ReportDes        string `json:"reportDes" gorm:"column:reportDes"`
	Status           string `json:"status" gorm:"column:status"`
	UpdateAt         string `json:"updateAt" gorm:"column:updateAt"`
	UpdateBy         int    `json:"updateBy" gorm:"column:updateBy"`
	CreateAt         string `json:"createdAt" gorm:"column:createAt"`
	CreateBy         int    `json:"createBy" gorm:"column:createBy"`
	RoomNum          string `json:"roomNum" gorm:"column:roomNum"`
	BuildingId       int    `json:"buildingId" gorm:"column:buildingId"`
	SelectedDate     string `json:"selectedDate" gorm:"column:selectedDate"`
}
