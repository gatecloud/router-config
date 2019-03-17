package models

type CustomConfig map[string]interface{}

type RouterTemplate struct {
	Resources     []string `sql:"-"`
	Methods       []string `sql:"-"`
	Version       string
	ProxySchema   string
	ProxyPass     string
	ProxyVersion  string
	CustomeConfig CustomConfig
}
