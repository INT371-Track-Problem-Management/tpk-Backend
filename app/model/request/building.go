package request

type Building struct {
	BuildingId int ` json:"buildingId" gorm:"column:buildingId"`
}

type BuildingInsert struct {
	BuildingName string `json:"buildingName" gorm:"column:buildingName"`
	UpdateBy     int    `json:"createBy" gorm:"column:createBy"`
}

type BuildingDelete struct {
	BuildingId int ` json:"buildingId" gorm:"column:buildingId"`
}
