package entity

type ProfileMedia struct {
	Id        int    `json:"id" gorm:"column:id"`
	Url       string `json:"url" gorm:"column:url"`
	Email     string `json:"email" gorm:"column:email"`
	ConfigKey string `json:"configKey" gorm:"column:configKey"`
	CreateAt  string `json:"createAt" gorm:"column:createAt"`
	UpdateAt  string `json:"updateAt" gorm:"column:updateAt"`
	UpdateBy  int    `json:"updateBy" gorm:"column:updateBy"`
}
