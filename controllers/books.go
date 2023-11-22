package controllers

import (
	"bookstore/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FindBooks(c *gin.Context) {
	var books []models.Book
	// Find all books matching the given conditions
	models.DB.Find(&books)
	c.JSON(http.StatusOK, gin.H{
		"data": books,
	})
}
