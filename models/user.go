package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID          int64  `json:"id" gorm:"primary_key"`
	Code        string `json:"code" gorm:"unique"`
	Name        string `json:"name" gorm:"varchar(255)"`
	Email       string `json:"email" gorm:"varchar(255) uniqueIndex index:code_index"`
	Username    string `json:"username" gorm:"varchar(100) uniqueIndex index:username_index"`
	Password    string `json:"password" gorm:"varchar(255)"`
	LastToken   string
	updateToken bool
	CreatedAt   time.Time `json:"created_at" gorm:"timestamp autoCreateTime:true"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"timestamp autoUpdateTime:true"`
	DeletedAt   gorm.DeletedAt
}
