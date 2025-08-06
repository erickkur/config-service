package router

import "net/http"

type RouterInterface interface {
	Handle(method string, pattern string, handler http.Handler)
	ServeHTTP(w http.ResponseWriter, req *http.Request)
}
