package response

type Building struct {
	BuildingId   int    `json:"buildingId" gorm:"column:buildingId"`
	BuildingName string `json:"buildingName" gorm:"column:buildingName"`
	Address      string `json:"address" gorm:"column:address"`
	Phone        string `json:"phone" gorm:"column:phone"`
	Email        string `json:"email" gorm:"column:email"`
}

type AllBuilding struct {
	BuildingId   int    `json:"buildingId" gorm:"column:buildingId"`
	BuildingName string `json:"buildingName" gorm:"column:buildingName"`
}
