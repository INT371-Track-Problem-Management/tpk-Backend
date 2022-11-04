package response

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

type Year struct {
	Year []int `json:"year" gorm:"column:year"`
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
}
