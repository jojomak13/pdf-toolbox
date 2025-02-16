package core

import (
	"bytes"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type Uploader struct {
	session *session.Session
	client  *s3manager.Uploader
}

var UploaderClient Uploader

func NewUploader() {
	s := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(os.Getenv("S3_REGION")),
		Credentials: credentials.NewStaticCredentials(os.Getenv("S3_KEY"), os.Getenv("S3_SECRET"), ""),
	}))

	UploaderClient = Uploader{
		session: s,
		client:  s3manager.NewUploader(s),
	}
}

func (u *Uploader) Upload(content []byte, path string) (string, error) {
	out, err := u.client.Upload(&s3manager.UploadInput{
		Bucket: aws.String(os.Getenv("S3_BUCKET")),
		Key:    aws.String(path),
		Body:   bytes.NewReader(content),
		ACL:    aws.String("public-read"),
	})

	if err != nil {
		return "", err
	}

	return out.Location, nil
}
