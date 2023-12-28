package common

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetPayment(c *gin.Context) {
	id := c.Param("id")

	_, err := ConnectToDB()
	if err != nil {
		fmt.Printf("Error conecting to db, %d", err)
		InternalServerError(c)
		return
	}

	c.IndentedJSON(200, gin.H{
		"id": id,
	})
}
