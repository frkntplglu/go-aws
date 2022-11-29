package main

import (
	"log"

	"github.com/frkntplglu/go-aws/s3"
)

const (
	AWS_S3_REGION     = "" // bucket region
	AWS_S3_BUCKET     = "" // bucket name
	AWS_S3_KEY_ID     = "" // access key id
	AWS_S3_SECRET_KEY = "" // secret key
)

func main() {
	s3 := s3.NewS3Client(AWS_S3_REGION, AWS_S3_KEY_ID, AWS_S3_SECRET_KEY)

	err := s3.UploadFile("support.jpg", "supportins3bucket.jpg", AWS_S3_BUCKET)

	if err != nil {
		log.Fatal(err)
	}
}
