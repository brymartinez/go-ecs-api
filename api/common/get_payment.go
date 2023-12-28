package common

import (
	"github.com/gin-gonic/gin"
)

func GetPayment(c *gin.Context) {
	id := c.Param("id")
	c.IndentedJSON(200, gin.H{
		"id": id,
	})
}
