package api

import (
	"github.com/gin-gonic/gin"
	"pkg/model/user"
)

// curl -I "http://127.0.0.1:8089/"
func Register() {
	r := gin.Default()
	r.GET("/user/register", func(c *gin.Context) {

		userService := &user.UserSerevice{}
		userInfo, err := userService.Insert()

		c.JSON(200, gin.H{
			"message":   "pong",
			"user_info": userInfo,
			"err":       err,
		})
	})
	r.Run(":8089") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
