package api

import (
	"net/http"
	"orders/database"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Ping is a small service check
func healthCheck(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}

// Get will find an order
func get(c *gin.Context) {
	str := c.Param("id")
	id, err := strconv.Atoi(str)
	if err != nil {
		handleError(c, err)
		return
	}

	order, err := database.FindById(uint(id))
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, order)
}

// Post will create an order
func post(c *gin.Context) {

	var request OrderRequestBody

	if err := c.BindJSON(&request); err != nil {
		handleError(c, err)
		return
	}

	id, err := database.Create(request.Quantity)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"Order created": id,
	})
}
