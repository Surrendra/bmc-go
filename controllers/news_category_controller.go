package controllers

import (
	"BaliMediaCenter/helpers"
	"BaliMediaCenter/middlewares"
	"fmt"
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
	user, _ := middlewares.GetSessionUser(c)
	fmt.Println(user.)
	con.ResponseHelper.ResponseSuccess(c, nil, "Success", 200)
}
