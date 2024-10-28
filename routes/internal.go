package routes

import (
	"BaliMediaCenter/controllers"
	"BaliMediaCenter/helpers"
	"BaliMediaCenter/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupInternalRoute(routes *gin.Engine) {
	ResponseHelper := helpers.NewResponseHelper()
	ValidationHelper := helpers.NewValidationHelper()
	StorageHelper := helpers.NewStorageHelper()
	PublicHelper := helpers.NewPublicHelper()

	NewsCategoryController := controllers.NewNewsCategoryController(ResponseHelper, ValidationHelper, StorageHelper, PublicHelper)

	internalRoute := routes.Group("/internal")
	internalRoute.Use(middlewares.AuthMiddleware())
	{
		newsCategoryRoute := internalRoute.Group("/news_category")
		{
			newsCategoryRoute.GET("get_data", NewsCategoryController.GetData)
			newsCategoryRoute.GET("get_data_with_pagination", NewsCategoryController.GetDataWithPagination)
			newsCategoryRoute.POST("create", NewsCategoryController.Create)
			newsCategoryRoute.GET("find_by_code/:code", NewsCategoryController.FindByCode)
			newsCategoryRoute.PUT("update/:code", NewsCategoryController.Update)
			newsCategoryRoute.DELETE("delete/:code", NewsCategoryController.Delete)
		}
	}
}
