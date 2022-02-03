package controller

import (
	"fmt"

	models "github.com/galigaribaldi/PersonalBooks/pkg/models"
)

func (h handler) CreateBook() {
	book := models.Book{
		Id:     3,
		Title:  "Flores en el ático",
		Age:    "1971",
		Author: "Virginia Clews Andrews (V.C Andrews)",
	}
	h.DB.Create(&book)
}

func (h handler) SelectAllBooks() {
	var books []models.Book
	h.DB.Find(&books)
	fmt.Println(books)
}
func (h handler) SelectBooksById(id int) {
	var book models.Book
	if result := h.DB.First(&book, id); result.Error != nil {
		fmt.Println("Errorsr, No hay!")
	}
	fmt.Println(book)
}
func deleteBook() {

}
func updateBook() {

}
func insertBook() {

}
