package s3

import (
	"bytes"
	"context"
	"io/ioutil"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type S3Client struct {
	Client   *s3.S3
	Uploader *s3manager.Uploader
}

func NewS3Client(region, keyId, secretKey string) *S3Client {
	session, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(keyId, secretKey, ""),
	})
	if err != nil {
		log.Fatal(err)
	}

	s3 := s3.New(session)

	uploader := s3manager.NewUploader(session)

	return &S3Client{
		Client:   s3,
		Uploader: uploader,
	}
}

func (s *S3Client) UploadFile(src, dest, bucket string) error {

	file, err := ioutil.ReadFile(src)
	if err != nil {
		log.Fatal(err)
	}

	input := &s3manager.UploadInput{
		Bucket:      aws.String(bucket),      // bucket's name
		Key:         aws.String(dest),        // files destination location
		Body:        bytes.NewReader(file),   // content of the file
		ContentType: aws.String("image/jpg"), // content type
	}

	res, err := s.Uploader.UploadWithContext(context.Background(), input)
	log.Printf("res %+v\n", res)
	log.Printf("err %+v\n", err)

	return err
}
