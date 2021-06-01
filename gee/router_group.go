package gee

import (
	"log"
	"net/http"
	"path"
)

// RouterGroup wraps up the Route to extend group management,
// just like a prefix manager
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

// Use adds middlewares to the group
func (g *RouterGroup) Use(middlewares ...HandlerFunc)  {
	g.middlewares = append(g.middlewares, middlewares...)
}

func (g *RouterGroup) createStaticHandler(relativePath string, fs http.FileSystem) HandlerFunc {
	absolutePath := path.Join(g.prefix, relativePath)
	fileServer := http.StripPrefix(absolutePath, http.FileServer(fs))
	return func(c *Context) {
		file := c.Param("filepath")
		if _, err := fs.Open(file); err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		fileServer.ServeHTTP(c.Writer, c.Req)
	}
}

// serve static files
func (g *RouterGroup) Static(relativePath, root string) {
	handler := g.createStaticHandler(relativePath, http.Dir(root))
	urlPattern := path.Join(relativePath, "/*filepath")
	// Register GET handlers
	g.GET(urlPattern, handler)
}