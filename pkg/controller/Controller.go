package controller

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dbURL := "postgresql://nnckmgzasylosx:99b170dfe60bb4771b0ac9edf19939df386bc88cf1222aa7fe69ce6446a4a254@ec2-44-196-68-164.compute-1.amazonaws.com:5432/ddaui1ic2hdida"

	database, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	DB = database
}
