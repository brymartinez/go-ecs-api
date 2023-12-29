package main

import (
	"go-ecs-api/api/handlers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err.Error())
	}
	router := gin.Default()
	router.GET("/payment/:id", handlers.GetPayment)
	router.POST("/payment", handlers.CreatePayment)
	// router.PUT("/payment/:id", updatePayment)
	// router.PUT("/payment/:id/cancel", cancelPayment)
	router.Run("localhost:3000")
}
