package models

import (
	"bytes"
	"encoding/json"
	"regexp"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type Template struct {
	Model
	Resource     string `validate:"required"`
	Method       string `validate:"required"`
	Version      string `validate:"required"`
	ProxySchema  string `validate:"required"`
	ProxyPass    string `validate:"required"`
	ProxyVersion string `validate:"required"`
	CustomConfig string `validate:"required"`
	ProjectName  string `validate:"required"`
	TemplateName string `validate:"required"`
	URL          string
}

func (t *Template) UploadFile(uploader *s3manager.Uploader, domain string, b []byte) error {
	name := t.ProjectName + t.TemplateName + ".json"
	bucket := regexp.MustCompile("/{1}[a-zA-Z0-9-]+/{1}").
		FindString(domain)
	result, err := uploader.Upload(
		&s3manager.UploadInput{
			Bucket: aws.String(bucket[1:len(bucket)]),
			Key:    aws.String(name),
			Body:   bytes.NewReader(b),
		})
	if err != nil {
		return err
	}
	t.URL = result.Location
	return nil
}

func (t *Template) Convert2RouterTemplate() (RouterTemplate, error) {
	var customConfig CustomConfig
	if err := json.Unmarshal([]byte(t.CustomConfig), &customConfig); err != nil {
		return RouterTemplate{}, err
	}
	return RouterTemplate{
		Resources:     strings.Split(filterString(t.Resource), " "),
		Methods:       strings.Split(filterString(t.Method), " "),
		Version:       t.Version,
		ProxySchema:   t.ProxySchema,
		ProxyPass:     t.ProxyPass,
		ProxyVersion:  t.ProxyVersion,
		CustomeConfig: customConfig,
	}, nil
}

func filterString(src string) string {
	dst := strings.Replace(src, "\r\n", "", -1)
	dst = strings.Replace(dst, "\"", "", -1)
	dst = strings.Replace(dst, "\t", "", -1)
	dst = strings.Replace(dst, "\n", "", -1)
	dst = strings.Replace(dst, "\r", "", -1)
	dst = strings.Replace(dst, " ", "", -1)
	dst = strings.Replace(dst, ",", " ", -1)
	return dst
}
