package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleError(c *gin.Context, err error) {
	log.Fatal(err)
	c.JSON(http.StatusInternalServerError, gin.H{
		"Error": "Unexpected",
	})
}
