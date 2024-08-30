package book

import (
	"fmt"
	"projeto-go/api-rest/pkg/kafka"
	"projeto-go/api-rest/pkg/models"
	"strconv"

	"gorm.io/gorm"
)

func FindAllBooks(db *gorm.DB) ([]models.Book, error) {
	var books []models.Book
	result := db.Find(&books)
	return books, result.Error
}

func FindBookById(db *gorm.DB, id uint) (models.Book, error) {
	var book models.Book
	result := db.First(&book, id)

	return book, result.Error
}

func AddBook(db *gorm.DB, book models.Book) (models.Book, error) {
	result := db.Create(&book)
	message := fmt.Sprintf("Book added: %s by %s", book.Name, book.Author)
	if err := kafka.SendMessage("test-topic", message); err != nil {
		return book, err
	}
	return book, result.Error
}

func ModifyBook(db *gorm.DB, id uint, updatedBook models.Book) (models.Book, error) {
	var book models.Book
	if err := db.Model(&book).Where("id = ?", id).
		Updates(models.Book{
			Author: updatedBook.Author,
			Name:   updatedBook.Name,
			Price:  updatedBook.Price}).
		Error; err != nil {

		return book, err
	}

	if err := db.First(&book, id).Error; err != nil {
		return book, err
	}

	message := fmt.Sprintf("Book updated: %s by %s", book.Name, book.Author)
	if err := kafka.SendMessage("test-topic", message); err != nil {
		return book, err
	}

	return book, nil
}

func RemoveBook(db *gorm.DB, id uint) error {
	result := db.Delete(&models.Book{}, id)
	message := fmt.Sprintf("Book deleted: ID - %s", strconv.Itoa(int(id)))
	if err := kafka.SendMessage("test-topic", message); err != nil {
		return err
	}
	return result.Error
}
