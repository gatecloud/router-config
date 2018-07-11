package logic

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

type Templator interface {
	Save() error
	Load() (Templator, error)
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

	// name := reflect.TypeOf(*r).Name()
	// if name == "" {
	// 	return errors.New("fail to reflect router template name")
	// }
	return ioutil.WriteFile("groups/"+fileName+".json", body, 0600)

}

func (r *RouterTemplate) Load() (Templator, error) {

	return nil, nil
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
