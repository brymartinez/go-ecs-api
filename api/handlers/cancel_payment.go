package handlers

import (
	"fmt"
	"go-ecs-api/api/common"
	"go-ecs-api/api/model"

	"github.com/gin-gonic/gin"
	pg "github.com/go-pg/pg/v10"
)

func CancelPayment(c *gin.Context) {
	id := c.Param("id")

	db, err := common.ConnectToDB()
	if err != nil {
		fmt.Printf("Error connecting to db, %d", err)
		common.InternalServerError(c)
		return
	}

	var payment model.Payment
	status := "CANCELLED"

	result, err := db.Model(&model.Payment{}).Set("status = ?", status).Where("id = ?", id).Returning("*").Update(&payment)
	if err != nil {
		fmt.Printf("Error cancelling payment, %d", err)
		if err.Error() == pg.ErrNoRows.Error() {
			common.NotFoundError(c)
		} else {
			common.InternalServerError(c)
		}
		return
	}

	fmt.Printf("%d", result)

	c.IndentedJSON(200, payment)
}
