package crud

import (
	"img-service/src/db"
	"img-service/src/models"
	"log"
)

func CreateFile(newFile models.FileBase) models.File {
	var newItem models.File
	newItem.FileBase = newFile
	conn := db.GetConnection()
	sqlconn, err := conn.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlconn.Close()

	log.Println("creating file")

	result := conn.Create(&newItem)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return newItem
}
