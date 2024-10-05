package routes

import (
	"BaliMediaCenter/controllers"
	"BaliMediaCenter/helpers"
	"BaliMediaCenter/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupInternalRoute(routes *gin.Engine) {
	ResponseHelper := helpers.NewResponseHelper()
	NewsCategoryController := controllers.NewNewsCategoryController(ResponseHelper)

	internalRoute := routes.Group("/internal")
	internalRoute.Use(middlewares.AuthMiddleware())
	{
		newsCategoryRoute := internalRoute.Group("/news_category")
		{
			newsCategoryRoute.GET("get_data", NewsCategoryController.GetData)
		}
	}
}
