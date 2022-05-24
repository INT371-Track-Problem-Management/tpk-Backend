package request

type Dorm struct {
	DormId int ` json:"dormId" gorm:"column:dormId"`
}

type DormInsert struct {
	Address string `json:"address" gorm:"column:address"`
	Phone   string `json:"phone" gorm:"column:phone"`
	Email   string `json:"email" gorm:"column:email"`
}

type DormDelete struct {
	DormId int ` json:"dormId" gorm:"column:dormId"`
}
