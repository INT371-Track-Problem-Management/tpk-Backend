package entity

type Test struct {
	UserId string `gorm:"column:USERID" json:"userId"`
	Name   string `gorm:"column:NAME" json:"name"`
}
