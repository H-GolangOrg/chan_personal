package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name     string  `gorm:"not null;column:user_name" json:"user_name"`
	Email    string  `gorm:"not null;column:user_email" json:"user_email"`
	Password string  `gorm:"not null;column:user_password" json:"user_password"`
	Boards   []Board `gorm:"foreignKey:UserId;references:UserId"`
	Replys   []Reply `gorm:"foreignKey:UserId;references:UserId"`
}
