package routes

import (
	"BaliMediaCenter/controllers"
	"BaliMediaCenter/helpers"
	"github.com/gin-gonic/gin"
)

func SetupAuthenticationRoute(routes *gin.Engine) {
	ResponHelper := helpers.NewResponseHelper()
	LoginController := controllers.NewLoginController(ResponHelper)

	authenticationRoute := routes.Group("/authentication")
	{
		authenticationRoute.POST("login", LoginController.Login)
	}
}
