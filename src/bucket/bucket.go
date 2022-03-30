package bucket

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const BUCKET = "img-bucket-service"

func GetClient() *s3.Client {

	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	// Create an Amazon S3 service client
	return s3.NewFromConfig(cfg)
}

func GetListObjects(client *s3.Client) map[string]int64 {
	output, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String(BUCKET),
	})
	if err != nil {
		log.Fatal(err)
	}

	list := make(map[string]int64)
	for _, object := range output.Contents {
		list[aws.ToString(object.Key)] = object.Size
		//log.Printf("key=%s size=%d", aws.ToString(object.Key), object.Size)
	}
	return list
}

func UploadObject(filename string, client *s3.Client, file *os.File) string {
	uploader := manager.NewUploader(client)
	result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(BUCKET),
		Key:    aws.String(filename),
		Body:   file,
	})
	if err != nil {
		log.Fatalf("Error uploading file: %s \n ", err.Error())
		return ""
	}
	return result.UploadID
}

func DownloadObject(filename string, client *s3.Client) int64 {
	file, err := os.Create("download.pdf")
	if err != nil {
		log.Fatalf("Error creating file: %s", err.Error())
		return 0
	}

	defer file.Close()

	obj := s3.GetObjectInput{
		Bucket: aws.String(BUCKET),
		Key:    aws.String(filename),
	}

	downloader := manager.NewDownloader(client)
	numBytes, err := downloader.Download(context.TODO(), file, &obj)

	if err != nil {
		log.Fatalf("Error downloading a file: %s", err.Error())
		return 0
	}

	return numBytes
}
