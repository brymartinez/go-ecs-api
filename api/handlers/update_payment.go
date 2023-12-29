package handlers

import (
	"fmt"
	"go-ecs-api/api/common"
	"go-ecs-api/api/model"

	"github.com/gin-gonic/gin"
)

func UpdatePayment(c *gin.Context) {
	id := c.Param("id")

	db, err := common.ConnectToDB()
	if err != nil {
		fmt.Printf("Error connecting to db, %d", err)
		common.InternalServerError(c)
		return
	}

	var payment model.Payment
	err = db.Model(&model.Payment{}).Where("id = ?", id).Select(&payment)
	if err != nil {
		fmt.Printf("Error getting payment, %d", err)
		if err.Error() == "pg: no rows in result set" {
			common.NotFoundError(c)
		} else {
			common.InternalServerError(c)
		}
		return
	}

	c.IndentedJSON(200, payment)
}
