package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	fmt.Print("Hello world")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusAccepted, gin.H{
			"message": "first server func",
		})
	})
	r.GET("/steps/:id", func(context *gin.Context) {
		steps := context.Param("id")
		context.JSON(http.StatusAccepted, steps)
	})

	r.Run()
}
