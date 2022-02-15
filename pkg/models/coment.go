package models

/*
	Tables created for make a different coments with Authors or Books
*/
//Coment for Book
type CommentBook struct {
	Id      int    `json: "id" gorm:"primaryKey;AUTO_INCREMENT"`
	Comment string `json:"comment" gorm:"default:'No Comments'`
	BookID  int    `json:"bookid"`
	Book    Book   `json:"-" gorm:"foreignKey:BookID;references:Id"`
}

//Coment for Author
type CommentAuthor struct {
	Id       int    `json: "id" gorm:"primaryKey;AUTO_INCREMENT"`
	Comment  string `json:"comment" gorm:"default:'No Comments'`
	AuthorID int    `json:"authorid"`
	Book     Author `json:"-" gorm:"foreignKey:AuthorID;references:Id"`
}
