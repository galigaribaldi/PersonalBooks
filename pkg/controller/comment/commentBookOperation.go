package comment

import (
	"fmt"
	"strconv"

	con "github.com/galigaribaldi/PersonalBooks/pkg/controller"
	models "github.com/galigaribaldi/PersonalBooks/pkg/models"
)

func init() {
	con.ConnectDatabase()
	con.DB.AutoMigrate()
	con.DB.AutoMigrate(&models.CommentBook{})
	con.DB.AutoMigrate(&models.CalificationBook{})
}

func CreateCommentBook(commentBook models.CommentBook) {
	con.DB.Create(&commentBook)
}

//Obtener Calificacion del Libro por ID del libro
func SelectAllCalificationByIdBook(idBook int) []models.CalificationBook {
	idString := strconv.Itoa(idBook)
	var newCalification []models.CalificationBook
	if result := con.DB.Where("book_id = ?", idString).Find(&newCalification); result.Error != nil {
		return newCalification
	}
	return newCalification
}

//Obtener Calificacion del Libro por Nombre del libro
func SelectAllCalificationByNameBook(name string) []models.CalificationBook {
	var newCalification []models.CalificationBook
	if result := con.DB.Where("title LIKE ?", "%"+name+"%").Find(&newCalification); result.Error != nil {
		return newCalification
	}
	return newCalification
}

//Obtener Calificaciones por puntuación
func SelectAllCalificationByCalification(calificaction int) []models.CalificationBook {
	calificactionBook := strconv.Itoa(calificaction)
	var newCalification []models.CalificationBook
	if result := con.DB.Where("book_id = ?", calificactionBook).Find(&newCalification); result.Error != nil {
		return newCalification
	}
	return newCalification
}
func SelectAllCommentBook() []models.CommentBook {
	var commentBook []models.CommentBook
	con.DB.Find(&commentBook)
	return commentBook
}

func SelectByIdCommentBook(id int) models.CommentBook {
	var commentBook models.CommentBook
	if result := con.DB.First(&commentBook, id); result.Error != nil {
		return commentBook
	}
	return commentBook
}
func DeleteCommentBook(id int) int {
	var commentBook models.CommentBook
	if result := con.DB.Delete(&commentBook, id); result.Error != nil {
		return 0
	}
	return id
}

func UpdateCommentBook(commentBook models.CommentBook) {
	con.DB.Model(&commentBook).Updates(models.CommentBook{
		Comment:             commentBook.Comment,
		BookID:              commentBook.BookID,
		Personal_Experience: commentBook.Personal_Experience,
	})
}

//Get Comment by Book ID
func SelectCommentsByBookId(id int) models.CommentBook {
	idString := strconv.Itoa(id)
	var newCommentBook models.CommentBook
	if result := con.DB.Where("book_id = ?", idString).Find(&newCommentBook); result.Error != nil {
		fmt.Println("error")
		return newCommentBook
	}
	return newCommentBook
}

//Get Book and CComment By Id Book
func SelectBookAndCommentByBookId(id int) ([]models.CommentBook, models.Book) {
	var newcommentBook []models.CommentBook
	var newBook models.Book
	if result := con.DB.Where("book_id = ?", strconv.Itoa(id)).Find(&newcommentBook); result.Error != nil {
		return newcommentBook, newBook
	}
	if result := con.DB.First(&newBook, id); result.Error != nil {
		return newcommentBook, newBook
	}
	return newcommentBook, newBook
}

//Get Book and CComment By Name Book
func SelectBookAndCommentByBookName(name string) ([]models.CommentBook, []models.Book) {
	var newcommentBook []models.CommentBook
	var newBook []models.Book
	if result := con.DB.Where("title LIKE ?", "%"+name+"%").Find(&newBook); result.Error != nil {
		return newcommentBook, newBook
	}
	if len(newBook) > 0 {
		con.DB.Find(&newcommentBook, newBook[0].Id)
	}

	return newcommentBook, newBook
}

/*
++++++++++++++++++++++++++++++++++++++++++++
			Calification Book
++++++++++++++++++++++++++++++++++++++++++++
*/
func CreateCalificationBook(calificationBook models.CalificationBook) {
	con.DB.Create(&calificationBook)
}

func UpdateCalificationBook(calificationBook models.CalificationBook) {
	con.DB.Model(&calificationBook).Updates(models.CalificationBook{
		CalificationNumber: calificationBook.CalificationNumber,
		CalificationWord:   calificationBook.CalificationWord,
		BookID:             calificationBook.BookID,
	})
}

//Get Book, Comment and Calification By Id Book
func SelectBookCommentCalificationByBookId(id int) ([]models.CommentBook, models.Book, []models.CalificationBook) {
	var newcommentBook []models.CommentBook
	var newBook models.Book
	var newCalification []models.CalificationBook
	if result := con.DB.Find(&newcommentBook, id); result.Error != nil {
		return newcommentBook, newBook, newCalification
	}
	if result := con.DB.Find(&newBook, id); result.Error != nil {
		return newcommentBook, newBook, newCalification
	}
	if result := con.DB.First(&newBook, id); result.Error != nil {
		return newcommentBook, newBook, newCalification
	}
	return newcommentBook, newBook, newCalification
}

//Get Book and Comment and Calification By Name Book
func SelectBookCommentCalificationByBookName(name string) ([]models.CommentBook, []models.Book) {
	var newcommentBook []models.CommentBook
	var newBook []models.Book
	if result := con.DB.Where("title LIKE ?", "%"+name+"%").Find(&newBook); result.Error != nil {
		return newcommentBook, newBook
	}
	if len(newBook) > 0 {
		con.DB.Find(&newcommentBook, newBook[0].Id)
	}

	return newcommentBook, newBook
}
