package handlers

import (
	"fmt"
	"go-ecs-api/api/common"

	"github.com/gin-gonic/gin"
)

func GetPayment(c *gin.Context) {
	id := c.Param("id")

	_, err := common.ConnectToDB()
	if err != nil {
		fmt.Printf("Error conecting to db, %d", err)
		common.InternalServerError(c)
		return
	}

	c.IndentedJSON(200, gin.H{
		"id": id,
	})
}
