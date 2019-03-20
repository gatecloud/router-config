package models

import (
	"bytes"
	"regexp"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// File
type File struct {
	Model
	Name      string     `validate:"required"`
	Templates []Template `validate:"gt=0" gorm:"many2many:file_templates"`
	URL       string
}

func (f *File) UploadFile(sess *session.Session, domain string, b []byte) error {
	name := f.Name + ".json"
	bucket := regexp.MustCompile("/{1}[a-zA-Z0-9-]+/{1}").
		FindString(domain)
	uploader := s3manager.NewUploader(sess)
	result, err := uploader.Upload(
		&s3manager.UploadInput{
			Bucket: aws.String(bucket[1:len(bucket)]),
			Key:    aws.String(name),
			Body:   bytes.NewReader(b),
		})
	if err != nil {
		return err
	}
	f.URL = result.Location
	return nil
}

func (f *File) DeleteFile(sess *session.Session, domain string) error {
	key := f.Name + ".json"
	bucket := regexp.MustCompile("/{1}[a-zA-Z0-9-]+/{1}").FindString(domain)
	input := &s3.DeleteObjectInput{
		Bucket: aws.String(bucket[1:len(bucket)]),
		Key:    aws.String(key),
	}
	svc := s3.New(sess)
	_, err := svc.DeleteObject(input)
	if err != nil {
		return err
	}
	return nil
}
