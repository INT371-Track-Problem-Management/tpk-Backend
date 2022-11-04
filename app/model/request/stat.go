package request

type Stat struct {
	Month int `json:"month" gorm:"column:month"`
	Year  int `json:"year" gorm:"column:year"`
}
