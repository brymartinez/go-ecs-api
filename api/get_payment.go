package paymentapi

import "github.com/gin-gonic/gin"

func getPayment(c *gin.Context) string {
	id := c.Param("id")
	return id
}
