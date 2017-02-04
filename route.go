package backend

import (
	"net/http"
	"src/samples"
)

func init() {
	// net/http
	http.HandleFunc("/01", samples.HelloWorld)
}
