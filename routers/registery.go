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
		libRoute.Route{
			Name:       "Files",
			Controller: reflect.TypeOf(controllers.FileController{}),
			Model:      &models.File{},
		},
		libRoute.Route{
			Name:       "Login",
			Controller: reflect.TypeOf(controllers.LoginController{}),
			Model:      &models.User{},
		},
		libRoute.Route{
			Name:       "Callback",
			Controller: reflect.TypeOf(controllers.CallbackController{}),
			Model:      &models.User{},
		},
	}
}
