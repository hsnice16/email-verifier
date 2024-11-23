package www

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()

	router.Use(cors.Default())
	routes(router)

	router.Run()
}

func routes(router *gin.Engine) {
	router.GET("/health", healthCheck)
}

func healthCheck(c *gin.Context) {
	c.Status(200)
}
