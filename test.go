package main

import (
	"fmt"
	//"log"

	book "github.com/galigaribaldi/PersonalBooks/pkg/controller/book"
	//con "github.com/galigaribaldi/PersonalBooks/pkg/controller"
	"github.com/galigaribaldi/PersonalBooks/pkg/models"
)

func main() {
	var newBook2 []models.Book
	newBook := models.Book{
		Id:     1,
		Author: "asda",
		Age:    "2013",
		Title:  "Book",
	}
	//Seleccionar todos
	newBook2 = book.SelectAllBooks()
	fmt.Println("Select All", newBook2)
	//Creadores de libro
	//book.CreateBook(newBook)
	//fmt.Println("Crear un libro:", newBook)
	//Update
	book.UpdateBook(newBook)
	//Seleccionar un libro
	fmt.Println(book.SelectBooksById(1))
	//Eliminar libro
	//book.DeleteBook(12)
	//fmt.Println(book.SelectAllBooks())

}
