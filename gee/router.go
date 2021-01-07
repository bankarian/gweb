package gee

import (
	"net/http"
	"strings"
)

// pattern is the path-model that register to routers, while
// path is the real time url-path
type router struct {
	// method : *node
	roots map[string]*node
	// method-pattern : handler
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

// addRoute registers a url pattern to routers
func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern)
	key := method + "-" + pattern
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = handler
}

// getRoute parses the real time url-path, then
// return the node in trie as well as the path-parameters
func (r *router) getRoute(method, path string) (*node, map[string]string) {
	realTimeParts := parsePattern(path)
	params := make(map[string]string)
	root, ok := r.roots[method]

	if !ok {
		return nil, nil
	}

	n := root.search(realTimeParts, 0)

	if n != nil {
		patternParts := parsePattern(n.pattern)
		// there may be 2 kinds of dynamic url
		// 1) /.../:param
		// 2) /.../*srcPath
		for i, p := range patternParts {
			if p[0] == ':' { // 1)
				params[p[1:]] = realTimeParts[i]
			} else if p[0] == '*' && len(p) > 1 { // 2)
				params[p[1:]] = strings.Join(realTimeParts[i:], "/")
				break
			}
		}
		return n, params
	}
	return nil, nil
}

// getRoutes gets all routes that binds the method
func (r *router) getRoutes(method string) []*node {
	root, ok := r.roots[method]
	if !ok {
		return nil
	}
	nodes := make([]*node, 0)
	root.travel(&nodes)
	return nodes
}

func (r *router) handle(c *Context) {
	n, params := r.getRoute(c.Method, c.Path)
	if n != nil {
		c.Params = params // update context
		key := c.Method + "-" + n.pattern
		r.handlers[key](c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}

// parsePattern divides the url pattern into multiple parts,
// allowing only one *
func parsePattern(pattern string) []string {
	ptns := strings.Split(pattern, "/")
	parts := make([]string, 0)
	for _, elm := range ptns {
		if elm != "" {
			parts = append(parts, elm)
			if elm[0] == '*' {
				break
			}
		}
	}
	return parts
}
