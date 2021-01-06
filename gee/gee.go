package gee

import (
	"log"
	"net/http"
)

// HandlerFunc defines the request handler used bu Gee
type HandlerFunc func(*Context)

// Engine implements the interface of ServeHTTP
type Engine struct {
	router *router
}

// New is the constructor of gee.Engine
func New() *Engine {
	return &Engine{router: newRouter()}
}

func (e *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	e.router.addRoute(method, pattern, handler)
}

// GET adds GET request
func (e *Engine) GET(path string, handler HandlerFunc) {
	e.addRoute("GET", path, handler)
}

// POST defines the method to add POST request
func (e *Engine) POST(path string, handler HandlerFunc) {
	e.addRoute("POST", path, handler)
}

// Run starts a http server
func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	e.router.handle(c)
}

