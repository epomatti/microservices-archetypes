package api

import (
	"github.com/gin-gonic/gin"
)

func Start() {
	r := gin.Default()

	r.GET("/health", healthCheck)
	r.GET("/api/orders/:id", get)
	r.POST("/api/orders", post)

	r.Run()
}
