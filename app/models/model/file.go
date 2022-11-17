package model

type ReportMedia struct {
	Id          string `json:"id" gorm:"column:id"`
	FileName    string `json:"file_name" gorm:"column:file_name"`
	Size        int64  `json:"size" gorm:"column:size"`
	ContentType string `json:"content_type" gorm:"column:content_type"`
	CreateAt    string `json:"createAt" gorm:"column:createAt"`
}

type ProfileMedia struct {
	Id          string `json:"id" gorm:"column:id"`
	FileName    string `json:"file_name" gorm:"column:file_name"`
	Size        int64  `json:"size" gorm:"column:size"`
	ContentType string `json:"content_type" gorm:"column:content_type"`
	CreateAt    string `json:"createAt" gorm:"column:createAt"`
	Email       string `gorm:"column:email" json:"email"`
}
