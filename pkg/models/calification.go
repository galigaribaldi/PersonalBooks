package models

type CalificationBook struct {
	ID                 int     `json: "id" gorm:"primaryKey;AUTO_INCREMENT"`
	CalificationNumber float32 `json: "calification_number"`
	CalificationWord   string  `json: "calification_word"`
	BookID             int     `json:"bookid"`
	Book               Book    `json:"-" gorm:"foreignKey:BookID;references:Id"`
}

type CalificationAuthor struct {
	ID                 int     `json: "id" gorm:"primaryKey;AUTO_INCREMENT"`
	CalificationNumber float32 `json: "calification_number"`
	CalificationWord   string  `json: "calification_word"`
	AuthorID           int     `json:"authorid"`
	Author             Author  `json:"-" gorm:"foreignKey:AuthorID;references:Id"`
}
