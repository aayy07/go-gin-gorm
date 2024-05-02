package main

import (
	"net/http"

	"github.com/aayy07/go-gin-gorm/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvvariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
