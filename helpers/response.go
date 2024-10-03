package helpers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type responseHelper struct {
}

func NewResponseHelper() *responseHelper {
	return &responseHelper{}
}

type ResponseHelper interface {
	ResponseSuccess(c *gin.Context, data interface{}, message string, code int)
}

func (h responseHelper) ResponseSuccess(c *gin.Context, data interface{}, message string, code int) {
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"data":    data,
		"message": message,
	})
}

func (h responseHelper) ResponseBadRequest(c *gin.Context, data interface{}, message string, code int) {
	c.JSON(http.StatusBadRequest, gin.H{
		"code":    code,
		"data":    data,
		"message": message,
	})
}
