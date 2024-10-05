package controllers

import (
	"BaliMediaCenter/helpers"
	"BaliMediaCenter/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type loginController struct {
	ResponseHelper helpers.ResponseHelper
	UserService    services.UserService
}

func NewLoginController(
	ResponseHelper helpers.ResponseHelper,
	UserService services.UserService,
) *loginController {
	return &loginController{
		ResponseHelper: ResponseHelper,
		UserService:    UserService,
	}
}

func (con loginController) Login(c *gin.Context) {
	res, errLogin, msg := con.UserService.Login(c.PostForm("username"), c.PostForm("password"))
	if errLogin != nil {
		con.ResponseHelper.ResponseBadRequest(c, errLogin, msg)
		return
	}
	con.ResponseHelper.ResponseSuccess(c, res, "Success", http.StatusOK)
}
