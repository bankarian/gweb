package main

import (
	"gitee.com/bankarian/gee-web/gee"
	"net/http"
)

func main() {
	e := gee.NewDefault()
	e.GET("/", func(c *gee.Context) {
		c.String(http.StatusOK, "Hello Beney")
	})
	e.GET("/panic", func(c *gee.Context) {
		names := []string{"wowow", "lala"}
		c.String(http.StatusOK, names[100])
	})
	g := e.Group("/group")
	{
		g.GET("/test", func(c *gee.Context) {
			c.String(http.StatusOK, "group succeed")
		})
	}
	e.Run(":9999")
}
