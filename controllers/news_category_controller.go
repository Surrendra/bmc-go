package controllers

import (
	"BaliMediaCenter/helpers"
	"BaliMediaCenter/models"
	"github.com/gin-gonic/gin"
)

type newsCategoryController struct {
	ResponseHelper helpers.ResponseHelper
}

func NewNewsCategoryController(ResponseHelper helpers.ResponseHelper) *newsCategoryController {
	return &newsCategoryController{
		ResponseHelper: ResponseHelper,
	}
}

func (con newsCategoryController) GetData(c *gin.Context) {
	// get news category
	NewsCategories := []models.NewsCategory{}
	models.DB.Preload("CreatedUser").Find(&NewsCategories)

	res := []models.NewsCategoryResponse{}
	for _, newsCategory := range NewsCategories {
		res = append(res, newsCategory.ToResponse())
	}
	con.ResponseHelper.ResponseSuccess(c, res, "Success", 200)
}

func (con newsCategoryController) Create(c *gin.Context) {

}
