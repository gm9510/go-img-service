package main

import (
	"img-service/src/bucket"
	"log"
	"os"
)

func main() {

	// Get the first page of results for ListObjectsV2 for a bucket

	client := bucket.GetClient()

	/*
		file, err := os.Open("Pago_25_nov_bcs.pdf")
		defer file.Close()

		if err != nil {
			log.Fatalf("Error opennig file: %s \n ", err.Error())
			os.Exit(1)
		}

		fileID := bucket.UploadObject("Pago_25_nov_bcs.pdf", client, file)
	*/

	list := bucket.GetListObjects(client)

	log.Println("first page results:")
	log.Println("%v", list)

	size := bucket.DownloadObject("Pago_25_nov_bcs.pdf", client)
	if size == 0 {
		log.Fatalf("Error while downloading")
		os.Exit(1)
	}

}
