package route

import (
	"net/http"
	"testing"

	. "../../jasmine"
)

func Test_(t *testing.T) {
	Describe("REST Route", func() {
		It("should be able to route a request", func() {
		})
		Describe("when the transfer is GET /loc/artical/:num", func() {
			r := &Rest{"GET", trimSplitPath("/loc/artical/:num"), nil}
			It("should match GET /loc/artical/105", func() {
				req, _ := http.NewRequest("GET", "/loc/artical/123", nil)
				Expect(r.IsMatch(req)).ToBeTruthy()
			})
			It("should not match GET /loc/uuu/artical/123", func() {
				req, _ := http.NewRequest("GET", "/loc/uuu/artical/123", nil)
				Expect(r.IsMatch(req)).ToBeFalsy()
			})
			It("should not match GET /loc/artical/123/rr", func() {
				req, _ := http.NewRequest("GET", "/loc/uuu/artical/123", nil)
				Expect(r.IsMatch(req)).ToBeFalsy()
			})
			It("should not match GET /artical/123", func() {
				req, _ := http.NewRequest("GET", "/artical/123", nil)
				Expect(r.IsMatch(req)).ToBeFalsy()
			})
			It("should not match POST /loc/artical/105", func() {
				req, _ := http.NewRequest("POST", "/loc/artical/005", nil)
				Expect(r.IsMatch(req)).ToBeFalsy()
			})
			Describe("when client request GET /loc/artical/123", func() {
				It("should add query :num=123", func() {
					req, _ := http.NewRequest("GET", "/loc/artical/123", nil)
					r.AddQueryString(req)
					Expect(req.URL.Query().Get(":num")).ToEqual("123")
				})
			})
		})
	})
}
