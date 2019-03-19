package models

import "fmt"

type CustomConfig map[string]interface{}

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
	fmt.Println(rt.Methods)
	for _, resource := range rt.Resources {
		for _, method := range rt.Methods {
			path := ""
			if rt.Version != "" {
				path = PATH_SEPARATOR + rt.Version + PATH_SEPARATOR + rt.RouterGroup + PATH_SEPARATOR + resource
			} else {
				path = PATH_SEPARATOR + rt.RouterGroup + PATH_SEPARATOR + resource
			}

			routers = append(routers, Router{
				Method:        method,
				Path:          path,
				ProxyScheme:   rt.ProxySchema,
				ProxyPass:     rt.ProxyPass,
				ProxyPassPath: PATH_SEPARATOR + resource,
				APIVersion:    rt.ProxyVersion,
				CustomConfigs: rt.CustomeConfig,
			})
		}
	}
	return routers
}
