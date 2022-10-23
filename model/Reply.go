package model

import "github.com/jinzhu/gorm"

type Reply struct {
	gorm.Model
	Content string `gorm:"not null;column:reply_content" json:"reply_content"`
	UserId  uint
	BoardId uint
}
