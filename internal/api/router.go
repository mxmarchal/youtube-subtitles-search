package api

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/health", GetHealthHandler)
	router.GET("/transcript/:videoId", GetTranscriptHandler)

	// Serve static files
	router.Static("/static", "./static")
	
	// Serve the index.html file
	router.GET("/", func(c *gin.Context) {
		c.File("./static/index.html")
	})

	return router
}