package routes

import (
	"BaliMediaCenter/controllers"
	"BaliMediaCenter/helpers"
	"BaliMediaCenter/services"
	"github.com/gin-gonic/gin"
)

func SetupAuthenticationRoute(routes *gin.Engine) {
	ResponHelper := helpers.NewResponseHelper()
	UserServices := services.NewUserService()
	LoginController := controllers.NewLoginController(ResponHelper, UserServices)

	authenticationRoute := routes.Group("/authentication")
	{
		authenticationRoute.POST("login", LoginController.Login)
	}
}
