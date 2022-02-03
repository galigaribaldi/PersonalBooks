package main

import (
	"log"

	controller "github.com/galigaribaldi/PersonalBooks/pkg/controller"
)

func main() {
	log.Println("Start Project")
	DB := controller.Init()
	h := controller.New(DB)
	//h.CreateBook()
	//h.SelectAllBooks()
	h.SelectBooksById(3)
	h.SelectBooksById(1)
	h.SelectBooksById(2)
}
