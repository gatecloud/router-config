package models

type CustomConfig map[string]interface{}

// RouterTemplate is the template of the router in a same group
type RouterTemplate struct {
	Resources     []string
	Methods       []string
	Version       string
	ProxySchema   string
	ProxyPass     string
	ProxyVersion  string
	CustomeConfig CustomConfig
	RouterGroup   string
}

const (
	PATH_SEPARATOR = "/" //string(os.PathSeparator)
)

func (rt *RouterTemplate) GenerateRouters() []Router {
	var (
		routers []Router
	)
	for _, resource := range rt.Resources {
		for _, method := range rt.Methods {
			var (
				path, proxyPassPath string
			)
			if rt.Version != "" {
				path = PATH_SEPARATOR + rt.Version
			}

			if rt.RouterGroup != "" {
				path += PATH_SEPARATOR + rt.RouterGroup
				proxyPassPath += PATH_SEPARATOR + rt.RouterGroup
			}

			proxyPassPath += PATH_SEPARATOR + resource
			path += PATH_SEPARATOR + resource

			routers = append(routers, Router{
				Method:        method,
				Path:          path,
				ProxyScheme:   rt.ProxySchema,
				ProxyPass:     rt.ProxyPass,
				ProxyPassPath: proxyPassPath,
				APIVersion:    rt.ProxyVersion,
				CustomConfigs: rt.CustomeConfig,
			})
		}
	}
	return routers
}
