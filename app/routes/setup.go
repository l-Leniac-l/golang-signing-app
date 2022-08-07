package routes

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
	router := gin.Default()

	identity := router.Group("/identities")
	{
		identity.POST("/sign", SignDocumentHandler)
	}
	router.GET("/health", HealthCheckHandler)

	return router
}
