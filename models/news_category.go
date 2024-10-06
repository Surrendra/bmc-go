package models

import "os"

type NewsCategory struct {
	Id              int64  `json:"id" gorm:"primary_key"`
	Code            string `json:"code" gorm:"unique"`
	Slug            string `json:"slug" gorm:"unique"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	CreatedUserId   int64  `json:"created_user_id"`
	CreatedUserName string `json:"created_user_name"`
	Image           string `json:"image"`
	ImageURL        string `json:"image_url"`
	CreatedUser     User   `gorm:"foreignkey:CreatedUserId;references:ID"`
}

func (u NewsCategory) GetImageURL() string {
	if u.Image == "" {
		return ""
	}
	return os.Getenv("STORAGE_URL") + u.Image
}

type NewsCategoryResponse struct {
	Id              int64  `json:"id"`
	Code            string `json:"code"`
	Slug            string `json:"slug"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	CreatedUserId   int64  `json:"created_user_id"`
	CreatedUserName string `json:"created_user_name"`
	Image           string `json:"image"`
	ImageURL        string `json:"image_url"`
}

func (u NewsCategory) ToResponse() NewsCategoryResponse {
	return NewsCategoryResponse{
		Id:              u.Id,
		Code:            u.Code,
		Slug:            u.Slug,
		Name:            u.Name,
		Description:     u.Description,
		CreatedUserId:   u.CreatedUserId,
		CreatedUserName: u.CreatedUserName,
		ImageURL:        u.GetImageURL(), // Populate the full image URL
	}
}
