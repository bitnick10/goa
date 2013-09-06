package route

import (
	"net/http"
	"strings"
)

type File struct {
	New string
	Old string
}

func (route *File) IsMatch(req *http.Request) bool {
	path := req.URL.Path
	return strings.HasPrefix(path, route.Old)
}
