package controllers

import (
	"github.com/Renewdxin/bookstore/models"
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

func CreateBook(c *gin.Context) {
	var input models.CreateBookInput
	//Validate input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//Create the book
	book := models.Book{Title: input.Title, Author: input.Author}
	models.DB.Create(&book)
	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}
func FindBook(c *gin.Context) {
	var book models.Book

	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": book})
}
