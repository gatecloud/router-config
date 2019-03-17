package routers

import (
	"reflect"
	"router-config/controllers"
	"router-config/models"

	libRoute "github.com/gatecloud/webservice-library/route"
)

var (
	RouteMap map[string][]libRoute.Route
)

func init() {
	RouteMap = make(map[string][]libRoute.Route)
	RouteMap["api"] = []libRoute.Route{
		libRoute.Route{
			Name:       "Projects",
			Controller: reflect.TypeOf(controllers.ProjectController{}),
			Model:      &models.Project{},
		},
		libRoute.Route{
			Name:       "Templates",
			Controller: reflect.TypeOf(controllers.TemplateController{}),
			Model:      &models.Template{},
		},
	}
}
