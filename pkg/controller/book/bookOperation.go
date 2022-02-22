package book

import (
	con "github.com/galigaribaldi/PersonalBooks/pkg/controller"
	models "github.com/galigaribaldi/PersonalBooks/pkg/models"
)

func init() {
	con.ConnectDatabase()
	con.DB.AutoMigrate()
	con.DB.AutoMigrate(&models.Book{})
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

func SelectBookByName(name string) []models.Book {
	var newBook []models.Book
	if result := con.DB.Where("title LIKE ?", "%"+name+"%").Find(&newBook); result.Error != nil {
		return newBook
	}
	return newBook
}

func UpdateBook(book models.Book) {
	con.DB.Model(&book).Updates(models.Book{
		Title:            book.Title,
		Author:           book.Author,
		Editorial:        book.Editorial,
		Saga:             book.Saga,
		Numero_Paginas:   book.Numero_Paginas,
		Anio_Publicacion: book.Anio_Publicacion,
		Link_portada:     book.Link_portada,
		Sinopsis:         book.Sinopsis,
	})
}
