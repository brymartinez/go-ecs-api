package main

import (
	"go-ecs-api/api/common"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type payment struct {
	ID        string    `json:"id"`
	Status    string    `json:"status"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err.Error())
	}
	router := gin.Default()
	router.GET("/payment/:id", common.GetPayment)
	// router.POST("/payment", createPayment)
	// router.PUT("/payment/:id", updatePayment)
	// router.PUT("/payment/:id/cancel", cancelPayment)
	router.Run("localhost:3000")
}
