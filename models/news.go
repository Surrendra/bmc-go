package models

import (
	"time"

	"gorm.io/gorm"
)

type News struct {
	ID             int64     `json:"id" gorm:"primary_key"`
	NewsCategoryId int64     `json:"news_category_id" gorm:"index"`
	CreatedUserId  int64     `json:"created_user_id" gorm:"index"`
	Code           string    `json:"code" gorm:"unique"`
	Title          string    `json:"title" gorm:"text"`
	Slug           string    `json:"slug" gorm:"unique"`
	Content        string    `json:"content" gorm:"text"`
	Image          string    `json:"image" gorm:"varchar(20)"`
	CommentType    string    `json:"comment_type" gorm:"varchar(20)"`
	Status         string    `json:"status" gorm:"varchar(20)"`
	AllowShare     string    `json:"allow_share" gorm:"char(4)"`
	CreatedAt      time.Time `json:"created_at" gorm:"timestamp autoCreateTime:true"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"timestamp autoUpdateTime:true"`
	CreatedUser    User      `gorm:"foreignkey:CreatedUserId;references:ID"`
	DeletedAt      gorm.DeletedAt
}
