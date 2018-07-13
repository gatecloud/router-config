package logic

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Templator interface {
	Save() error
	Load(filename string) error
	Parse(args ...string) error
}

type CustomConfig map[string]interface{}

type RouterTemplate struct {
	Resources     []string
	Methods       []string
	Version       string
	ProxySchema   string
	ProxyPass     string
	ProxyVersion  string
	CustomConfigs CustomConfig
}

func (r *RouterTemplate) Save(fileName string) error {
	body, err := json.MarshalIndent(r, "", "\t")
	if err != nil {
		return err
	}

	return ioutil.WriteFile("groups/"+fileName+".json", body, 0600)
}

func (r *RouterTemplate) Parse(resources, methods, version, proxySchema, proxyPass, proxyVersion, customConfig string) error {
	resources = filterString(resources)
	r.Resources = strings.Split(resources, " ")
	methods = filterString(methods)
	methods = strings.ToUpper(methods)
	r.Methods = strings.Split(methods, " ")
	r.Version = version
	r.ProxySchema = proxySchema
	r.ProxyPass = proxyPass
	r.ProxyVersion = proxyVersion
	return json.Unmarshal([]byte(customConfig), &r.CustomConfigs)
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

func (r *RouterTemplate) Load(filename string) error {
	if filename == "" {
		return errors.New("filename can not be empty")
	}

	if err := filepath.Walk("groups/", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			_, fname := filepath.Split(path)
			if fname == filename {
				body, err := ioutil.ReadFile(path)
				if err != nil {
					return err
				}

				err = json.Unmarshal(body, r)
				if err != nil {
					return err
				}

				return nil
			}
		}
		return nil
	}); err != nil {
		return nil
	}
	return nil
}
