package logic

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

type Templator interface {
	Save() error
	Load() (Templator, error)
}

type CustomConfig map[string]interface{}

type RouterTemplate struct {
	Resources     []string
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

func (r *RouterTemplate) Parse(resources, version, proxySchema, proxyPass, proxyVersion, customConfig string) error {
	resources = strings.Replace(resources, "\r\n", "", -1)
	resources = strings.Replace(resources, "\"", "", -1)
	resources = strings.Replace(resources, "\t", "", -1)
	resources = strings.Replace(resources, "\n", "", -1)
	resources = strings.Replace(resources, "\r", "", -1)
	resources = strings.Replace(resources, " ", "", -1)
	r.Resources = strings.Split(resources[:len(resources)-1], ",")
	r.Version = version
	r.ProxySchema = proxySchema
	r.ProxyPass = proxyPass
	r.ProxyVersion = proxyVersion
	return json.Unmarshal([]byte(customConfig), &r.CustomConfigs)
}
