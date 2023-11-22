package main

import (
	"bookstore/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()
	// To define a route, we need the endpoint and the handler
	// func(context *gin.Context){}: determines how we provide the data to the client
	r.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"data": "Hello World",
		})
	})
	r.Run(":9090")
}
