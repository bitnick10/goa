package express

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/Bitnick2002/goa/express/route"
)

type RouteCollection struct {
	restRoutes []route.Rest
	fileRoutes []route.File
	Static     string
}

func New() *RouteCollection {
	return &RouteCollection{}
}
func used() {
	fmt.Println("...")
}

func (route *RouteCollection) Get(params string, handler http.HandlerFunc) {
	route.AddRoute("GET", params, handler)
}
func (route *RouteCollection) Post(params string, handler http.HandlerFunc) {
	route.AddRoute("POST", params, handler)
}
func (route *RouteCollection) Put(params string, handler http.HandlerFunc) {
	route.AddRoute("POST", params, handler)
}
func (route *RouteCollection) Delete(params string, handler http.HandlerFunc) {
	route.AddRoute("POST", params, handler)
}
func (route *RouteCollection) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	for _, rest := range route.restRoutes {
		if rest.IsMatch(req) {
			rest.AddQueryString(req)
			rest.Handler(res, req)
			return
		}
	}
	for _, fileRoute := range route.fileRoutes {
		if fileRoute.IsMatch(req) {
			wd, _ := os.Getwd()
			fullPath := wd + fileRoute.New + strings.TrimPrefix(req.URL.Path, fileRoute.Old)
			http.ServeFile(res, req, fullPath)
			return
		}
	}
	wd, _ := os.Getwd()
	fullPath := wd + route.Static + req.URL.Path
	http.ServeFile(res, req, fullPath)
}

// func (route *RouteCollection) Static(params, dir string) {
// 	route.AddRoute("GET", params, func(res http.ResponseWriter, req *http.Request) {
// 		http.ServeFile(res, req, params)
// 	})
// }
func (collection *RouteCollection) AddStatic(static, old string) {
	collection.fileRoutes = append(collection.fileRoutes, route.File{static, old})
}
func (collection *RouteCollection) AddRoute(method string, params string, handler http.HandlerFunc) {
	trimParams := strings.Trim(params, "/")
	locations := strings.Split(trimParams, "/")
	collection.restRoutes = append(collection.restRoutes, route.Rest{method, locations, handler})
}
