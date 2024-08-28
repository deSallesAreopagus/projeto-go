package router

import (
	"projeto-go/pkg/book"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	bookRoutes := r.Group("/books")
	{
		bookRoutes.GET("/", book.GetBooks(db))
		bookRoutes.GET("/:id", book.GetBookById(db))
		bookRoutes.POST("/", book.CreateBook(db))
		bookRoutes.PUT("/:id", book.UpdateBook(db))
		bookRoutes.DELETE("/:id", book.DeleteBook(db))
	}

	return r
}
