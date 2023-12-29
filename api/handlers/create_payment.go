package handlers

import (
	"fmt"
	"go-ecs-api/api/common"

	"github.com/gin-gonic/gin"
)

func CreatePayment(c *gin.Context) {

	db, err := common.ConnectToDB()
	if err != nil {
		fmt.Printf("Error conecting to db, %d", err)
		common.InternalServerError(c)
		return
	}

	_, err := db.Model(payment).Insert(payment)
	if err != nil {
		common.InternalServerError(c)
		return
	}

	c.IndentedJSON(200, payment)
}
