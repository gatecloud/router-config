package models

// Router is server configuration structure
type Router struct {
	Method        string
	Path          string
	ProxyScheme   string
	ProxyPass     string
	ProxyPassPath string
	APIVersion    string
	CustomConfigs CustomConfig
}