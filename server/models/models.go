package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID 	 uint   `json:"id" gorm:"primaryKey"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Media struct {
	gorm.Model
	ID       uint   `json:"id" gorm:"primaryKey"`
	MediaType string `json:"mediaType"`
	Title    string `json:"title"`
	Author    string `json:"author"`
	Likes	[]uint `json:"likes" gorm:"type:integer[]"`
	Favorites	[]uint `json:"favorites" gorm:"type:integer[]"`
	Url      string `json:"url"`
	ImgUrl   string `json:"imgUrl"`
	CreatorID uint `json:"creatorID"`
	CreatedAt string `json:"created_at"`
}
