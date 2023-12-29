package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InternalServerError(c *gin.Context) {
	c.IndentedJSON(http.StatusInternalServerError, gin.H{
		"code":    9999,
		"message": "Internal Server Error",
	})
}

func NotFoundError(c *gin.Context) {
	c.IndentedJSON(http.StatusNotFound, gin.H{
		"code":    1000,
		"message": "Not found",
	})
}

func InvalidStatusError(c *gin.Context) {
	c.IndentedJSON(http.StatusBadRequest, gin.H{
		"code":    1001,
		"message": "Invalid status",
	})
}
