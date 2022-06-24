package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

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

	timestamp := time.Now().Unix()

	dst := fmt.Sprint("./temp/file_%v.pdf", timestamp)
	c.SaveUploadedFile(file, dst)

	var newFile models.FileBase
	newFile.Name = fmt.Sprintf("%v_%s", timestamp, file.Filename)
	newFile.Size = uint(file.Size)
	newFile.Ext = file.Header.Get("Content-Type")

	// open file and upload to aws
	toUpload, _ := os.Open(dst)

	client := bucket.GetClient()
	UploadID := bucket.UploadObject(newFile.Name, client, toUpload)

	log.Printf("type %T %T\n", toUpload, UploadID)

	//close and remove
	err = toUpload.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = os.Remove(dst)
	if err != nil {
		log.Fatal(err)
	}

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
