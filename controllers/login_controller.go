package controllers

import (
	"BaliMediaCenter/helpers"
	"github.com/gin-gonic/gin"
	"net/http"
)

type loginController struct {
	ResponseHelper helpers.ResponseHelper
}

func NewLoginController(
	ResponseHelper helpers.ResponseHelper,
) *loginController {
	return &loginController{
		ResponseHelper: ResponseHelper,
	}
}

func (controller loginController) Login(c *gin.Context) {
	controller.ResponseHelper.ResponseSuccess(c, nil, "Success", http.StatusOK)
}
