package main

import (
	"baotian0506.com/39_config/gocode/mygin/start_some_service"
	"baotian0506.com/39_config/gocode/mygin/start_some_service/mysky"
	"github.com/gin-gonic/gin"
)

func main() {
	start_some_service.Start()
	mysky.Start()
	return
	some()
	return
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func loginEndpoint(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

var submitEndpoint = loginEndpoint
var readEndpoint = loginEndpoint

func some() {
	router := gin.Default()

	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.GET("/login", loginEndpoint)
		v1.POST("/submit", submitEndpoint)
		v1.POST("/read", readEndpoint)
	}

	// Simple group: v2
	v2 := router.Group("/v2")
	{

		v2.POST("/login", loginEndpoint)
		v2.POST("/submit", submitEndpoint)
		v2.POST("/read", readEndpoint)
	}

	router.Run(":8080")

}

func some2() {
	gin.Mode()
}
