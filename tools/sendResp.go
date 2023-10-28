package tools

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleError(c *gin.Context, err error, message string) bool {
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": message + ": " + err.Error(),
		})
		return true
	}
	return false
}

func SendResp(c *gin.Context, code int, T any, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"data":    T,
		"message": message,
	})
}

func HandleErrOrResp(c *gin.Context, resp any, err error) {
	if HandleError(c, err, "") {
		return
	}
	SendResp(c, http.StatusOK, resp, "")
}
