package router

import (
	"net/http"

	"new-token/pkg/log"
)

// HandlerNotFound Function
func handlerNotFound(w http.ResponseWriter, r *http.Request) {
	log.Println(log.LogLevelWarn, "http-access", "not found method "+r.Method+" at URI "+r.RequestURI)
	ResponseNotFound(w, "not found method "+r.Method+" at URI "+r.RequestURI)
}

// HandlerMethodNotAllowed Function
func handlerMethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	log.Println(log.LogLevelWarn, "http-access", "not allowed method "+r.Method+" at URI "+r.RequestURI)
	ResponseMethodNotAllowed(w, "not allowed method "+r.Method+" at URI "+r.RequestURI)
}
