package gee

import "log"

type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc // support middleware
	parent      *RouterGroup  // support nesting
	engine      *Engine       // all groups share an Engine instance
}

func (g *RouterGroup) Group(prefix string) *RouterGroup {
	engine := g.engine // get the shared engine
	newGroup := &RouterGroup{
		prefix: g.prefix + prefix,
		parent: g,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (g *RouterGroup) addRoute(method string, pattern string, handler HandlerFunc) {
	wholePattern := g.prefix + pattern
	log.Printf("Route %4s - %s", method, wholePattern)
	g.engine.router.addRoute(method, wholePattern, handler)
}

// GET adds GET request
func (g *RouterGroup) GET(path string, handler HandlerFunc) {
	g.addRoute("GET", path, handler)
}

// POST defines the method to add POST request
func (g *RouterGroup) POST(path string, handler HandlerFunc) {
	g.addRoute("POST", path, handler)
}
