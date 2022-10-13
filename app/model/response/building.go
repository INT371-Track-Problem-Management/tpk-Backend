package response

type Building struct {
	BuildingId 		int    `json:"buildingId" gorm:"column:BUILDINGID"`
	BuildingName 	string `json:"buildingName" gorm:"column:BUILDINGNAME"`
	Address    		string `json:"address" gorm:"column:ADDRESS"`
	Phone      		string `json:"phone" gorm:"column:PHONE"`
	Email      		string `json:"email" gorm:"column:EMAIL"`
}
