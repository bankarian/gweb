package gee

import (
	"fmt"
	"reflect"
	"testing"
)

func newTestRouter() *router {
	r := newRouter()
	r.addRoute("GET", "/", nil)
	r.addRoute("GET", "/greet/people/employee/:name", nil)
	r.addRoute("GET", "/greet/people/client/:name", nil)
	r.addRoute("GET", "/assets/*filepath", nil)
	return r
}

func TestParsePattern(t *testing.T) {
	ok := reflect.DeepEqual(parsePattern("/p/:name"), []string{"p", ":name"})
	ok = ok && reflect.DeepEqual(parsePattern("/p/*"), []string{"p", "*"})

	// allows only one *
	ok = ok && reflect.DeepEqual(parsePattern("/p/*name/*"), []string{"p", "*name"})
	if !ok {
		t.Fatal("test parsePattern failed")
	}
}



func TestGetRoutes(t *testing.T) {
	r := newTestRouter()
	nodes := r.getRoutes("GET")
	for i, n := range nodes {
		fmt.Println(i+1, n)
	}
	if len(nodes) != 3 {
		t.Fatal("the number of routes should be 3")
	}
}

func BenchmarkGetRoute(b *testing.B) {
	r := newTestRouter()
	// for i := 0; i < b.N; i++ {
	// 	r.getRoute("GET", "/greet/people/student")
	// }
	for i := 0; i < b.N; i++ {
		r.getRoute("GET", "/greet/people/client/beney")
	}
}
