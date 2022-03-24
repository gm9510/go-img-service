package main

import (
	"img-service/src/bucket"
	"log"
)

func main() {

	// Get the first page of results for ListObjectsV2 for a bucket

	client := bucket.GetClient()
	list := bucket.GetListObjects(client)

	log.Println("first page results:")
	log.Println("%v", list)
}
