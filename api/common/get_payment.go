package common

import (
	"github.com/gin-gonic/gin"
)

func GetPayment(c *gin.Context) {
	id := c.Param("id")

	_, err := ConnectToDB()
	if err != nil {
		InternalServerError(c)
		return
	}

	c.IndentedJSON(200, gin.H{
		"id": id,
	})
}
