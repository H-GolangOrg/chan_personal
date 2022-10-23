package model

import "github.com/jinzhu/gorm"

type Board struct {
	gorm.Model
	Title   string  `gorm:"not null;column:board_title" json:"board_title"`
	Content string  `gorm:"not null;column:board_content" json:"board_content"`
	Replys  []Reply `gorm:"foreignKey:BoardId;references:BoardId"`
	UserId  uint
}
