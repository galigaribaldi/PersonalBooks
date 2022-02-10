package book

import (
	con "github.com/galigaribaldi/PersonalBooks/pkg/controller"
	models "github.com/galigaribaldi/PersonalBooks/pkg/models"
)

func init() {
	con.ConnectDatabase()
}
func CreateBook(book models.Book) {
	con.DB.Create(&book)
}

func SelectAllBooks() []models.Book {
	var books []models.Book
	con.DB.Find(&books)
	return books
}

func SelectBooksById(id int) models.Book {
	var book models.Book
	if result := con.DB.First(&book, id); result.Error != nil {
		return book
	}
	return book
}
func DeleteBook(id int) {
	var book models.Book
	if result := con.DB.Delete(&book, id); result.Error != nil {
		return
	}
}

func UpdateBook(book models.Book) {
	con.DB.Model(&book).Updates(models.Book{
		Title:  book.Title,
		Age:    book.Age,
		Author: book.Author,
	})
}