package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

type payment struct {
	ID        string    `json:"id"`
	Status    string    `json:"status"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func main() {
	router := gin.Default()
	router.GET("/payment/:id", getPayment)
	router.POST("/payment", createPayment)
	router.PUT("/payment/:id", updatePayment)
	router.PUT("/payment/:id/cancel", cancelPayment)
	router.Run("localhost:3000")
}
