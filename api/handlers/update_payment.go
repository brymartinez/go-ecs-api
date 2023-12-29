package handlers

import (
	"fmt"
	"go-ecs-api/api/common"
	"go-ecs-api/api/model"
	"slices"

	"github.com/gin-gonic/gin"
	pg "github.com/go-pg/pg/v10"
)

var validState = []string{"CREATED"}

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
		if err.Error() == pg.ErrNoRows.Error() {
			common.NotFoundError(c)
		} else {
			common.InternalServerError(c)
		}
		return
	}

	if !slices.Contains(validState, payment.Status) {
		common.InvalidStatusError(c)
		return
	}

	db.Model(&payment).Set("status = ?", "COMPLETED").Returning("*").Where("id = ?", id).Update(&payment)
	if err != nil {
		fmt.Printf("Error cancelling payment, %d", err)
		if err.Error() == pg.ErrNoRows.Error() {
			common.NotFoundError(c)
		} else {
			common.InternalServerError(c)
		}
		return
	}

	c.IndentedJSON(200, payment)
}
