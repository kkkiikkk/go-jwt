package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string  `gorm:"uniqueIndex;not null" json:"username"`
	Password string  `gorm:"not null" json:"password"`
	Images   []Image `json:"images`
}

type Image struct {
	gorm.Model
	ImagePath string `gorm:"uniqueIndex;not null" json:"image_path"`
	ImageUrl  string `gorm:"uniqueIndex;not null" json:"image_url"`
	UserId 	  uint	 `gorm:"not null" json:"user_id"`
}