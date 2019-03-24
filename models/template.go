package models

import (
	"encoding/json"
	"strings"
)

type Template struct {
	Model
	Resource     string `validate:"required"`
	Method       string `validate:"required"`
	Version      string
	ProxySchema  string `validate:"required"`
	ProxyPass    string `validate:"required"`
	ProxyVersion string
	CustomConfig string `validate:"required"`
	ProjectName  string `validate:"required"`
	RouterGroup  string `validate:"required"`
	TemplateName string `validate:"required"`
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
		RouterGroup:   t.RouterGroup,
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
