package NotFound

import (
	"net/http"
)

func H(res http.ResponseWriter, req *http.Request) {
	http.Error(res, "NOT FOUND", 404)
}
