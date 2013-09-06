package express

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	. "../jasmine"
)

func Test_ServeHTTP(t *testing.T) {
	req, _ := http.NewRequest("GET", "/loc1/jeson/artical/125", nil)
	res := httptest.NewRecorder()

	handler := new(RouteCollection)
	handler.Get("/loc1/:user/artical/:num", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "hello world")
		res.WriteHeader(http.StatusOK)
	})
	handler.ServeHTTP(res, req)

	// fmt.Println("raw query = ", req.URL.RawQuery)

	user := req.URL.Query().Get(":user")
	num := req.URL.Query().Get(":num")

	Expect(user).ToEqual("jeson")
	Expect(num).ToEqual("125")
}

func Test_Static(t *testing.T) {
	Describe("route static ", func() {
		route := new(RouteCollection)
		route.Static = "/web"
		It("should get default static file", func() {
			req, _ := http.NewRequest("GET", "/for_static_test.html", nil)
			w := httptest.NewRecorder()
			route.ServeHTTP(w, req)
			pwd, _ := os.Getwd()

			testFile, _ := ioutil.ReadFile(pwd + "/web/for_static_test.html")
			Expect(w.Body.String()).ToEqual(string(testFile))
		})
		It("should get custom static file", func() {
			route.AddStatic("/route", "/a/b/c")
			req, _ := http.NewRequest("GET", "/a/b/c/file.go", nil)
			w := httptest.NewRecorder()
			route.ServeHTTP(w, req)
			pwd, _ := os.Getwd()

			testFile, _ := ioutil.ReadFile(pwd + "/route/file.go")
			Expect(w.Body.String()).ToEqual(string(testFile))
		})
	})
}
