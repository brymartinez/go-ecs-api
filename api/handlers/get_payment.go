package handlers

import (
	"fmt"
	"go-ecs-api/api/common"
	model "go-ecs-api/api/models"

	"github.com/gin-gonic/gin"
)

func GetPayment(c *gin.Context) {
	id := c.Param("id")

	db, err := common.ConnectToDB()
	if err != nil {
		fmt.Printf("Error conecting to db, %d", err)
		common.InternalServerError(c)
		return
	}

	var payment model.Payment
	db.Model(&model.Payment{}).Where("id = ?", id).Select(&payment)

	c.IndentedJSON(200, payment)
}
