package request

type Maintainer struct {
	Fname string `json:"fname" gorm:"column:fname"`
	Lname string `json:"lname" gorm:"column:lname"`
	Phone string `json:"phone" gorm:"column:phone"`
}
