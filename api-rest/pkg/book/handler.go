package book

import (
	"net/http"
	"projeto-go/api-rest/pkg/models"

	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetBooks(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		books, err := FindAllBooks(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}
		c.JSON(http.StatusOK, books)
	}
}

func GetBookById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		bookID, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		book, err := FindBookById(db, uint(bookID))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "ID not found"})
			return
		}
		c.JSON(http.StatusOK, book)
	}
}

func CreateBook(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var book models.Book
		if err := c.BindJSON(&book); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid values"})
			return
		}

		createdBook, err := AddBook(db, book)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create book"})
			return
		}
		c.JSON(http.StatusCreated, createdBook)
	}
}

func UpdateBook(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		bookID, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		var updatedBook models.Book
		if err := c.BindJSON(&updatedBook); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid values"})
			return
		}

		book, err := ModifyBook(db, uint(bookID), updatedBook)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}
		c.JSON(http.StatusOK, book)
	}
}

func DeleteBook(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		bookID, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		err = RemoveBook(db, uint(bookID))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}
		c.JSON(http.StatusNoContent, nil)
	}
}
