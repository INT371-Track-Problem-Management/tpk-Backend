package request

type Dorm struct {
	DormId int ` json:"dormId" gorm:"column:dormId"`
}
