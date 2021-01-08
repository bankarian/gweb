package gee

import (
	"net/http"
)

// a map to store JSON k-v
type H map[string]interface{}

// HandlerFunc defines the request handler used by Gee
type HandlerFunc func(*Context)


// Engine implements the interface http.Handler
type Engine struct {
	*RouterGroup
	router *router
	groups []*RouterGroup // store all groups
}

// NewEngine is the constructor of gee.Engine
func NewEngine() *Engine {
	e := &Engine{router: newRouter()}
	e.RouterGroup = &RouterGroup{engine: e}
	e.groups = []*RouterGroup{e.RouterGroup}
	return e
}

// Run starts a http server
func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	e.router.handle(c)
}
