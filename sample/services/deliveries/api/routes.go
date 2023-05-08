package api

import (
	"github.com/gin-gonic/gin"
)

func Start() {
	r := gin.Default()

	r.GET("/health", healthCheck)
	r.GET("/api/deliveries/:id", get)
	r.POST("/api/deliveries", post)

	r.Run()
}
