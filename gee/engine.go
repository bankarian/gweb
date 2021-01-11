package gee

import (
	"html/template"
	"net/http"
	"strings"
)

// a map to store JSON k-v
type H map[string]interface{}

// HandlerFunc defines the request handler used by Gee
type HandlerFunc func(*Context)

// Engine implements the interface http.Handler
type Engine struct {
	*RouterGroup
	router        *router
	groups        []*RouterGroup     // store all groups
	htmlTemplates *template.Template // for html render
	funcMap       template.FuncMap   // for html render
}

// New is the constructor of gee.Engine
func New() *Engine {
	e := &Engine{router: newRouter()}
	e.RouterGroup = &RouterGroup{engine: e}
	e.groups = []*RouterGroup{e.RouterGroup}
	return e
}

// NewDefault uses Logger() & Recovery() middlewares
func NewDefault() *Engine {
	e := New()
	e.Use(Logger(), Recovery())
	return e
}

// Run starts a http server
func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}

// customize render function
func (e *Engine) SetFuncMap(funcMap template.FuncMap) {
	e.funcMap = funcMap
}

func (e *Engine) LoadHTMLGlob(pattern string) {
	e.htmlTemplates = template.
		Must(template.New("").
			Funcs(e.funcMap).
			ParseGlob(pattern))
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var middlewares []HandlerFunc
	for _, g := range e.groups {
		if strings.HasPrefix(req.URL.Path, g.prefix) {
			middlewares = append(middlewares, g.middlewares...)
		}
	}
	c := newContext(w, req)
	c.handlers = middlewares
	c.engine = e
	e.router.handle(c)
}
