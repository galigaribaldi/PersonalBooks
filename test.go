package main

import (

	//"log"

	author "github.com/galigaribaldi/PersonalBooks/pkg/controller/author"

	//con "github.com/galigaribaldi/PersonalBooks/pkg/controller"
	"github.com/galigaribaldi/PersonalBooks/pkg/models"
)

func main() {
	//var newBook2 []models.Book
	//var newAuthor2 []models.Author
	/*
		newBook := models.Book{
			Id:     1,
			Author: "asda",
			Age:    "2013",
			Title:  "Book",
		}
	*/
	//Seleccionar todos
	//newBook2 = book.SelectAllBooks()
	//fmt.Println("Select All Books", newBook2)
	///Create New Author
	newAuthor := models.Author{
		Nombre:           "Virginia Clew Andrews (VC Andrews)",
		Anio_Nacimiento:  1940,
		Anio_Defuncion:   1981,
		Pais_Origen:      "Estados Unidos de América (EUA)",
		Motivo_Defuncion: "Infarto",
		Photo_Img:        "https://upload.wikimedia.org/wikipedia/en/2/2f/V._C._Andrews.jpg",
		BookID:           1,
	}
	author.CreateAuthor(newAuthor)
	///
	//newAuthor2 = author.SelectAllAuthor()
	//fmt.Println("Select All Author", newAuthor2)
	//Creadores de libro
	//book.CreateBook(newBook)
	//fmt.Println("Crear un libro:", newBook)
	//Update
	//book.UpdateBook(newBook)
	//Seleccionar un libro
	//fmt.Println(book.SelectBooksById(1))
	//Eliminar libro
	//book.DeleteBook(12)
	//fmt.Println(book.SelectAllBooks())

}
