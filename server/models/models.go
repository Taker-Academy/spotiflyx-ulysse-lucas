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
	Url      string `json:"url"`
	ImgUrl   string `json:"imgUrl"`
	CreatorID uint `json:"creatorID"`
	CreatedAt string `json:"created_at"`
}

type Interaction struct {
	gorm.Model
	ID       uint   `json:"id" gorm:"primaryKey"`
	UserID   uint   `json:"userID"`
	MediaID  uint   `json:"mediaID"`
	InteractionType string `json:"interactionType"`
}
