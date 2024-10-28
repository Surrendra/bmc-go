package helpers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type responseHelper struct{}

func NewResponseHelper() *responseHelper {
	return &responseHelper{}
}

type ResponseHelper interface {
	ResponseSuccess(c *gin.Context, data interface{}, message string, code int)
	ResponseBadRequest(c *gin.Context, err error, message string)
	ResponseSuccessWithPagination(c *gin.Context, data interface{}, pageSize int, pageIndex int)
}

func (h responseHelper) ResponseSuccess(c *gin.Context, data interface{}, message string, code int) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": data,
		"msg":  message,
	})
}

func (h responseHelper) ResponseBadRequest(c *gin.Context, err error, message string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"code": http.StatusBadRequest,
		"err":  err,
		"msg":  message,
	})
}

func (h responseHelper) ResponseSuccessWithPagination(c *gin.Context, data interface{}, pageSize int, pageIndex int) {
	PublicHelper := NewPublicHelper()
	paginationResponse := map[string]string{}
	paginationResponse = PublicHelper.generatePaginationUrl(c, pageSize, pageIndex)
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": map[string]interface{}{
			"data":          data,
			"current_page":  paginationResponse["current_page"],
			"previous_page": paginationResponse["previus_page"],
		},
		"msg": "Success",
	})
}
