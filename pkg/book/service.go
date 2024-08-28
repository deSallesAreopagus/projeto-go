package book

import (
	"projeto-go/pkg/models"

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

	return book, nil
}

func RemoveBook(db *gorm.DB, id uint) error {
	result := db.Delete(&models.Book{}, id)
	return result.Error
}
