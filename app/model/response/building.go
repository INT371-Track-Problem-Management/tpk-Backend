package response

type Building struct {
	BuildingId int    ` json:"dormId" gorm:"column:DORMID"`
	Address    string `json:"address" gorm:"column:ADDRESS"`
	Phone      string `json:"phone" gorm:"column:PHONE"`
	Email      string `json:"email" gorm:"column:EMAIL"`
}
