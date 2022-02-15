package models

//Table for Author
type Author struct {
	Id               int    `json: "id" gorm:"primaryKey;AUTO_INCREMENT"`
	Nombre           string `json:"nombre"`
	Anio_Nacimiento  int    `json:"anio_nacimiento"`
	Anio_Defuncion   int    `json:"anio_defuncion"`
	Pais_Origen      string `json:"pais_origen"`
	Motivo_Defuncion string `json:"motivo_defuncion"`
	Photo_Img        string `json:"photo_img"`
	BookID           int    `json:"bookid"`
	Mini_Biografia   string `json:"mini_biografia"`
	Book             Book   `json:"-" gorm:"foreignKey:BookID;references:Id"`
}
