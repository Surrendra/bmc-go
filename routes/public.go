package routes

import (
	"BaliMediaCenter/helpers"
	"fmt"
	"github.com/gin-gonic/gin"
)

func SetupPublicRoute(route *gin.Engine) {
	route.GET("/file/:directory/:filename", func(c *gin.Context) {
		filename := c.Param("filename")
		directory := c.Param("directory")
		StorageHelper := helpers.NewStorageHelper()

		//filePath := filepath.Join("storages/app/", directory, "/", filename)
		filePath := StorageHelper.GetPullFilePath(directory, filename)
		fmt.Println(directory, filename, "open url", filePath)
		// Check if the file exists and serve it
		c.File(filePath)
	})
}
