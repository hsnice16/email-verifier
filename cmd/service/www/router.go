package www

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hsnice16/email-verifier/cmd/docs"
	"github.com/swaggo/files"       // swagger embed files
	"github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func Run() {
	router := gin.Default()

	router.Use(cors.Default())
	routes(router)

	docs.SwaggerInfo.Title = "Email Verifier Service Endpoint"
	docs.SwaggerInfo.Description = "This is a Email Verifier Service server."
	docs.SwaggerInfo.Version = "1.0"

	// use ginSwagger middleware to serve the API docs
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run()
}

func routes(router *gin.Engine) {
	router.GET("/health", healthCheck)
}

// healthCheck
//
//	@Summary	Health Check
//	@Tags		General
//	@Success	200
//	@Router		/health [get]
func healthCheck(c *gin.Context) {
	c.Status(200)
}
