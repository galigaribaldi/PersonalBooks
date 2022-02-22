package author

import (
	con "github.com/galigaribaldi/PersonalBooks/pkg/controller"
	models "github.com/galigaribaldi/PersonalBooks/pkg/models"
)

func init() {
	con.ConnectDatabase()
	con.DB.AutoMigrate()
	con.DB.AutoMigrate(&models.Author{})
}

func CreateAuthor(author models.Author) {
	con.DB.Create(&author)
}

func SelectAllAuthor() []models.Author {
	var authors []models.Author
	con.DB.Find(&authors)
	return authors
}

func SelectAuthorById(id int) models.Author {
	var author models.Author
	if result := con.DB.First(&author, id); result.Error != nil {
		return author
	}
	return author
}

func SelectAuthorByName(name string) []models.Author {
	var newAuthor []models.Author
	if result := con.DB.Where("nombre LIKE ?", "%"+name+"%").Find(&newAuthor); result.Error != nil {
		return newAuthor
	}
	return newAuthor
}

func DeleteAuthor(id int) {
	var author models.Author
	if result := con.DB.Delete(&author, id); result.Error != nil {
		return
	}
}

func UpdateAuthor(author models.Author) {
	con.DB.Model(&author).Updates(models.Author{
		Nombre:           author.Nombre,
		Anio_Nacimiento:  author.Anio_Nacimiento,
		Anio_Defuncion:   author.Anio_Defuncion,
		Pais_Origen:      author.Pais_Origen,
		Motivo_Defuncion: author.Motivo_Defuncion,
		Photo_Img:        author.Photo_Img,
		Mini_Biografia:   author.Mini_Biografia,
	})
}
