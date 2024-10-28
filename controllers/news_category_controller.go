package controllers

import (
	"BaliMediaCenter/helpers"
	"BaliMediaCenter/middlewares"
	"BaliMediaCenter/models"
	"BaliMediaCenter/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"strconv"
)

type newsCategoryController struct {
	ResponseHelper   helpers.ResponseHelper
	ValidationHelper helpers.ValidationHelper
	StorageHelper    helpers.StorageHelper
	PublicHelper     helpers.PublicHelper
}

func NewNewsCategoryController(
	ResponseHelper helpers.ResponseHelper,
	ValidationHelper helpers.ValidationHelper,
	StorageHelper helpers.StorageHelper,
	PublicHelper helpers.PublicHelper,
) *newsCategoryController {
	return &newsCategoryController{
		ResponseHelper:   ResponseHelper,
		ValidationHelper: ValidationHelper,
		StorageHelper:    StorageHelper,
		PublicHelper:     PublicHelper,
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
	newsCategory := models.NewsCategory{}
	err := c.Bind(&newsCategory)
	if err != nil {
		con.ResponseHelper.ResponseBadRequest(c, err, "")
		return
	}

	file, fileHeader, err := c.Request.FormFile("image")
	if err != nil {
		log.Println(err)
		con.ResponseHelper.ResponseBadRequest(c, err, "failed get image from request")
		return
	}

	errValidation := con.ValidationHelper.ValidateImage(file, fileHeader)
	if errValidation != nil {
		log.Println(errValidation)
		con.ResponseHelper.ResponseBadRequest(c, errValidation, "Opps Something wrong with the image")
		return
	}

	errUpload, filepath := con.StorageHelper.UploadCategoryNews(file, fileHeader)
	if errUpload != nil {
		con.ResponseHelper.ResponseBadRequest(c, errUpload, errUpload.Error())
		return
	}

	user, _ := middlewares.GetSessionUser(c)
	newsCategory.Name = c.PostForm("name")
	newsCategory.Description = c.PostForm("description")
	newsCategory.Slug = con.PublicHelper.MakeSlugFromString(newsCategory.Name)
	newsCategory.CreatedUserId = user.ID
	newsCategory.CreatedUserName = user.Name
	newsCategory.Image = filepath
	newsCategory.Code = uuid.NewString()
	NewsCategoryService := services.NewNewsCategoryService()
	NewsCategoryService.Create(newsCategory)

	newsCategoryResponse := models.NewsCategoryResponse{}
	newsCategoryResponse = NewsCategoryService.FindByCodeWithResponseFormat(newsCategory.Code)
	con.ResponseHelper.ResponseSuccess(c, newsCategoryResponse, "Success", http.StatusOK)
}

func (con newsCategoryController) FindByCode(c *gin.Context) {
	code := c.Param("code")
	newsCategory := models.NewsCategoryResponse{}
	NewsCategoryService := services.NewNewsCategoryService()
	newsCategory = NewsCategoryService.FindByCodeWithResponseFormat(code)
	if newsCategory.Code == "" {
		con.ResponseHelper.ResponseBadRequest(c, nil, "Product Category not found !")
		return
	}
	con.ResponseHelper.ResponseSuccess(c, newsCategory, "Success", http.StatusOK)
}

func (con newsCategoryController) Update(c *gin.Context) {
	NewsCategoryService := services.NewNewsCategoryService()
	code := c.Param("code")
	newsCategory := models.NewsCategory{}
	ExistingNewsCategory := models.NewsCategory{}
	ExistingNewsCategory = NewsCategoryService.FindByCode(code)
	fmt.Println(ExistingNewsCategory.Name)
	if ExistingNewsCategory.Code == "" {
		con.ResponseHelper.ResponseBadRequest(c, nil, "Product Category not found !")
		return
	}
	newsCategory.Image = ExistingNewsCategory.Image
	file, fileHeader, _ := c.Request.FormFile("image")
	if file != nil {
		errUpload, filepath := con.StorageHelper.UploadCategoryNews(file, fileHeader)
		if errUpload != nil {
			con.ResponseHelper.ResponseBadRequest(c, errUpload, "Opps Something wrong with the image")
			return
		}
		newsCategory.Image = filepath
	}
	newsCategory.Code = code
	newsCategory.Name = c.PostForm("name")
	newsCategory.Description = c.PostForm("description")
	NewsCategoryService.Update(newsCategory, code)
	newsCategoryResponse := models.NewsCategoryResponse{}
	newsCategoryResponse = NewsCategoryService.FindByCodeWithResponseFormat(newsCategory.Code)
	con.ResponseHelper.ResponseSuccess(c, newsCategoryResponse, "Successfully updated the record !", http.StatusOK)
}

func (con newsCategoryController) Delete(c *gin.Context) {
	code := c.Param("code")
	NewsCategoryService := services.NewNewsCategoryService()
	ExistingNewsCategory := models.NewsCategory{}
	ExistingNewsCategory = NewsCategoryService.FindByCode(code)
	if ExistingNewsCategory.Code == "" {
		con.ResponseHelper.ResponseBadRequest(c, nil, "Product Category not found !")
		return
	}
	NewsCategoryService.Delete(code)
	con.ResponseHelper.ResponseSuccess(c, nil, "Successfully deleted the record", http.StatusOK)
}

func (con newsCategoryController) GetDataWithPagination(c *gin.Context) {
	//fmt.Println("GetDataWithPagination : ", c.Param("page_index"))
	NewsCategoryService := services.NewNewsCategoryService()
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	pageIndex, _ := strconv.Atoi(c.Query("page_index"))
	if pageSize == 0 {
		pageSize = 10
	}
	if pageIndex == 0 {
		pageIndex = 1
	}
	//fmt.Println(pageSize, pageIndex)
	newsCategoryResponse := []models.NewsCategoryResponse{}
	newsCategoryResponse, errPaginate := NewsCategoryService.GetPaginateData(pageSize, pageIndex)
	if errPaginate != nil {
		con.ResponseHelper.ResponseBadRequest(c, nil, "Something wrong when get data")
		return
	}
	con.ResponseHelper.ResponseSuccessWithPagination(c, newsCategoryResponse, pageSize, pageIndex)
}
