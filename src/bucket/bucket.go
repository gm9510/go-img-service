package bucket

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const BUCKET = "employee-photo-bucket-al-137"

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
