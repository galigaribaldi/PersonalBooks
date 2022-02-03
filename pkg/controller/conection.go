package controller

import (
	"log"

	"github.com/galigaribaldi/PersonalBooks/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dbURL := "postgresql://nnckmgzasylosx:99b170dfe60bb4771b0ac9edf19939df386bc88cf1222aa7fe69ce6446a4a254@ec2-44-196-68-164.compute-1.amazonaws.com:5432/ddaui1ic2hdida"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(&models.Book{})
	return db
}
