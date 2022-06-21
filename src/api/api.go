package api

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"img-service/src/bucket"
	"img-service/src/crud"
	"img-service/src/models"

	"github.com/gin-gonic/gin"
)

func ReadAllFiles(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ReadAllFiles",
	})
}

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		log.Fatal(err)
	}

	c.SaveUploadedFile(file, "./temp/example.pdf")

	var newFile models.FileBase
	newFile.Name = file.Filename
	newFile.Size = uint(file.Size)
	newFile.Ext = file.Header.Get("Content-Type")

	toUpload, _ := os.Open("./temp/example.pdf")

	client := bucket.GetClient()
	UploadID := bucket.UploadObject(newFile.Name, client, toUpload)

	log.Printf("type %T %v\n", toUpload, UploadID)

	c.JSON(http.StatusOK, crud.CreateFile(newFile))
}

func DownloadById(c *gin.Context) {
	id := c.Param("id")
	msg := fmt.Sprintf("DownloadById %v", id)
	c.JSON(http.StatusOK, gin.H{
		"message": msg,
	})
}

func UpdateById(c *gin.Context) {
	id := c.Param("id")
	msg := fmt.Sprintf("UpdateById %v", id)
	c.JSON(http.StatusOK, gin.H{
		"message": msg,
	})
}

func DeleteById(c *gin.Context) {
	id := c.Param("id")
	msg := fmt.Sprintf("DeleteById %v", id)
	c.JSON(http.StatusOK, gin.H{
		"message": msg,
	})
}

func PartialUpdateById(c *gin.Context) {
	id := c.Param("id")
	msg := fmt.Sprintf("PartialUpdateById %v", id)
	c.JSON(http.StatusOK, gin.H{
		"message": msg,
	})
}
