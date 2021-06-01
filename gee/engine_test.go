package gee

import (
	"fmt"
	"testing"
)

func TestNestedGroup(t *testing.T) {
	e := New()
	v1 := e.Group("/v1")
	v2 := v1.Group("/v2")
	v3 := v2.Group("/v3")
	fmt.Println(v1.prefix, v2.prefix, v3.prefix)
}
