package main

import (
	"log"
	"net/http"

	"img-service/src/api"
	"img-service/src/db"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db.Migrate()

	r := gin.Default()

	files := r.Group("/files")
	{
		files.GET("", api.ReadAllFiles)
		files.POST("", api.UploadFile)
		files.GET(":id", api.DownloadById)
		files.PUT(":id", api.UpdateById)
		files.DELETE(":id", api.DeleteById)
		files.PATCH(":id", api.PartialUpdateById)
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
