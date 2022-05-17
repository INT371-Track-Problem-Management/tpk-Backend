package entity

type Dorm struct {
	DormId  int    `json:"dormId" gorm:"column:dormId"`
	Address string `json:"address" gorm:"column:address"`
	Phone   string `json:"phone" gorm:"column:phone"`
	Email   string `json:"email" gorm:"column:email"`
}
