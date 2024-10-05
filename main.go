package main

import (
	"BaliMediaCenter/databases/seeds"
	"BaliMediaCenter/models"
	"BaliMediaCenter/routes"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {
	fmt.Println("initiate system")
	models.ConnectDatabase()
	db := models.DB
	seeds.SeedAll(db)
	const AppProduction = "production"
	AppEnv := os.Getenv("APP_ENV")
	if AppEnv == AppProduction {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	route := gin.Default()
	route.GET("/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "Status ok",
		})
	})
	routes.SetupAuthenticationRoute(route)
	routes.SetupInternalRoute(route)
	route.Run()
}
