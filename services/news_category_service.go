package services

import (
	"BaliMediaCenter/models"
	"time"
)

type newsCategoryService struct {
}

func NewNewsCategoryService() *newsCategoryService {
	return &newsCategoryService{}
}

type NewsCategoryService interface {
	Create(NewsCategory models.NewsCategory) interface{}
	FindByCode(code string) interface{}
}

func (s newsCategoryService) Create(newsCategory models.NewsCategory) interface{} {
	models.DB.Create(&newsCategory)
	newsCategory.CreatedAt = time.Now()
	newsCategory.UpdatedAt = time.Now()
	models.DB.Preload("CreatedUser").First(&newsCategory)
	return newsCategory
}

func (s newsCategoryService) FindByCode(code string) models.NewsCategory {
	newsCategory := models.NewsCategory{}
	models.DB.Preload("CreatedUser").Where("code = ?", code).Find(&newsCategory)
	return newsCategory
}

func (s newsCategoryService) FindByCodeWithResponseFormat(code string) models.NewsCategoryResponse {
	newsCategoryResponse := models.NewsCategoryResponse{}
	newsCategory := models.NewsCategory{}
	models.DB.Preload("CreatedUser").Where("code = ?", code).Find(&newsCategory)
	newsCategoryResponse = newsCategory.ToResponse()
	return newsCategoryResponse
}

func (s newsCategoryService) Update(newsCategory models.NewsCategory, code string) interface{} {
	newsCategory.UpdatedAt = time.Now()
	models.DB.Model(&newsCategory).Where("code = ?", code).Updates(newsCategory)
	return newsCategory
}

func (s newsCategoryService) Delete(code string) {
	newsCategory := models.NewsCategory{}
	models.DB.Where("code = ?", code).Delete(&newsCategory)
}

func (s newsCategoryService) GetPaginateData(pageSize int, pageIndex int) ([]models.NewsCategoryResponse, error) {
	NewsCategories := []models.NewsCategory{}
	offset := (pageIndex - 1) * pageSize
	err := models.DB.Limit(pageSize).Offset(offset).Model(&NewsCategories).Find(&NewsCategories).Error
	if err != nil {
		return nil, err
	}
	res := []models.NewsCategoryResponse{}
	for _, newsCategory := range NewsCategories {
		res = append(res, newsCategory.ToResponse())
	}
	return res, nil
}
