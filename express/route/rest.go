package route

import (
	// "fmt"
	"net/http"
	"net/url"
	"strings"
)

type Rest struct {
	Method  string
	Params  []string
	Handler http.HandlerFunc
}

func trimSplitPath(path string) []string {
	trimPath := strings.Trim(path, "/")
	return strings.Split(trimPath, "/")
}

func (this *Rest) IsMatch(req *http.Request) bool {
	parts := trimSplitPath(req.URL.Path)
	if this.Method != req.Method {
		return false
	}
	if len(this.Params) != len(parts) {
		return false
	}
	for i := range this.Params {
		if this.Params[i][0] == ":"[0] || this.Params[i] == parts[i] {
			continue
		} else {
			return false
		}
	}
	return true
}

func (this *Rest) AddQueryString(req *http.Request) {
	parts := trimSplitPath(req.URL.Path)
	values := req.URL.Query()
	for i := range this.Params {
		if this.Params[i][0] == ":"[0] {
			// fmt.Println("adding ", this.Params[i], " ", parts[i])
			values.Add(this.Params[i], parts[i]) // req.URL.Query().Add(this.Params[i], parts[i])
		}
	}
	req.URL.RawQuery = url.Values(values).Encode() + "&" + req.URL.RawQuery
	// fmt.Println("after adding ")
	// fmt.Println(req.URL.Query().Get(":name"))
	// fmt.Println(req.URL.RawQuery)
}

// values := r.URL.Query()
// 			for i, match := range matches[1:] {
// 				values.Add(Rest.Params[i], match)
// 			}

// 			//reassemble query Params and add to RawQuery
// 			r.URL.RawQuery = url.Values(values).Encode() + "&" + r.URL.RawQuery
