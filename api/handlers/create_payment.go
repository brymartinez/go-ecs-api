package handlers

import (
	"fmt"
	"go-ecs-api/api/common"
	"go-ecs-api/api/model"
	"time"

	"github.com/gin-gonic/gin"
)

func CreatePayment(c *gin.Context) {
	var requestBody model.CreatePaymentDTO

	if err := c.BindJSON(&requestBody); err != nil {
		// DO SOMETHING WITH THE ERROR
		fmt.Printf("Error parsing body, %d", err)
		common.InternalServerError(c)
		return
	}

	fmt.Printf("%+v\n", requestBody)

	db, err := common.ConnectToDB()
	if err != nil {
		fmt.Printf("Error connecting to db, %d", err)
		common.InternalServerError(c)
		return
	}

	var payment model.Payment
	payment.Status = "CREATED"
	payment.Amount = requestBody.Amount
	payment.CreatedAt = time.Now()
	payment.UpdatedAt = time.Now()
	_, err = db.Model(&payment).Insert(&payment)
	if err != nil {
		fmt.Printf("Error saving to db, %d", err)
		common.InternalServerError(c)
		return
	}

	c.IndentedJSON(200, payment)
}
