package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Uid         string `gorm:"uid" json:"uid"`
	Email       string `gorm:"email" json:"email"`
	Password    string `gorm:"password" json:"password"`
	Fullname    string `gorm:"fullname" json:"fullname"`
	PhoneNumber string `gorm:"phone_number" json:"phone_number"`
}
