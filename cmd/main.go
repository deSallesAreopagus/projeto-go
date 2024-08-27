package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID     string  `json:"id"`
	Author string  `json:"author"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
}

var books = []book{
	{ID: "1", Author: "Autor 1", Name: "Livro 1", Price: 59.99},
	{ID: "2", Author: "Autor 2", Name: "Livro 2", Price: 79.99},
	{ID: "3", Author: "Autor 3", Name: "Livro 3", Price: 109.99},
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)

	router.Run("localhost:8000")
}

func getBooks(b *gin.Context) {
	b.JSON(http.StatusOK, books)
}
