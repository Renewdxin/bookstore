package main

import (
	"github.com/Renewdxin/bookstore/controllers"
	"github.com/Renewdxin/bookstore/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()
	// To define a route, we need the endpoint and the handler
	// func(context *gin.Context){}: determines how we provide the data to the client
	r.GET("/books", controllers.FindBooks)
	r.GET("/books/:id", controllers.FindBook)
	r.POST("/books", controllers.Createbook)
	r.Run(":9090")
}
