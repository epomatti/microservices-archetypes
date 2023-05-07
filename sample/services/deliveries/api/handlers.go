package api

import (
	"deliveries/database"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Ping is a small service check
func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// Get will find a delivery
func get(c *gin.Context) {
	str := c.Param("id")
	id, err := strconv.Atoi(str)
	if err != nil {
		handleError(c, err)
		return
	}

	d, err := database.FindById(uint(id))
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, d)
}

// Post will create a delivery
func post(c *gin.Context) {

	var request DeliveryRequestBody

	if err := c.BindJSON(&request); err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Unexpected",
		})
		return
	}

	id, err := database.Create(request.Address, request.OrderID)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Unexpected",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"Delivery created": id,
	})
}
