package models

import (
	"bytes"
	"encoding/json"
	"regexp"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	FILE_SIZE int = 2048
)

// File
type File struct {
	Model
	Name      string     `validate:"required"`
	Templates []Template `validate:"gt=0" gorm:"many2many:file_templates"`
	URL       string
	Preview   string `sql:"-"`
}

func (f *File) UploadFile(sess *session.Session, domain string, b []byte) error {
	name := f.Name + ".json"
	bucket := regexp.MustCompile("/{1}[a-zA-Z0-9-]+/{1}").
		FindString(domain)

	svc := s3.New(sess)
	input := &s3.PutObjectInput{
		Body:   aws.ReadSeekCloser(bytes.NewReader(b)),
		Bucket: aws.String(bucket[1:len(bucket)]),
		Key:    aws.String(name),
		Metadata: map[string]*string{
			"metadata1": aws.String("text/plain"),
			"metadata2": aws.String("application/json"),
		},
	}

	if _, err := svc.PutObject(input); err != nil {
		return err
	}

	f.URL = domain + name
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

func (f *File) GetFileContent(sess *session.Session, domain string) error {
	key := f.Name + ".json"
	bucket := regexp.MustCompile("/{1}[a-zA-Z0-9-]+/{1}").FindString(domain)
	input := &s3.GetObjectInput{
		Bucket: aws.String(bucket[1:len(bucket)]),
		Key:    aws.String(key),
	}
	svc := s3.New(sess)
	result, err := svc.GetObject(input)
	if err != nil {
		return err
	}

	b := make([]byte, FILE_SIZE, FILE_SIZE)
	_, err = result.Body.Read(b)
	if err != nil {
		return err
	}

	body, err := json.MarshalIndent(b, "", "\t")
	if err != nil {
		return err
	}

	f.Preview = string(body)
	return nil
}
