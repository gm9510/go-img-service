package db

import (
	"log"
	"os"

	//"log"
	"img-service/src/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetConnection() *gorm.DB {

	host := os.Getenv("HOST")
	//port     := os.Getenv("PORT")
	user := os.Getenv("DBUSER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DATABASE")

	connection_string := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=5432 sslmode=disable"
	db, error := gorm.Open(postgres.Open(connection_string), &gorm.Config{})

	if error != nil {
		log.Fatal(error)
	}

	return db
}

func Migrate() {
	db := GetConnection()
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	log.Println("Automigration working....")

	db.AutoMigrate(&models.File{})
}
