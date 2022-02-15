package models

type Book struct {
	Id               int    `json: "id" gorm:"primaryKey;AUTO_INCREMENT"`
	Title            string `json:"title"`
	Author           string `json: "author"`
	Editorial        string `json: "editorial"`
	Saga             string `json: "saga"`
	Numero_Paginas   int    `json: "numero_paginas"`
	Anio_Publicacion int    `json: "anio_publicacion"`
	Link_portada     string `json: "link_portada"`
	Sinopsis         string `json: "sinopsis"`
}
