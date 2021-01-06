package main

import (
	"net/http"

	"gitee.com/bankarian/gee-web/gee"
)

func main() {
	r := gee.New()
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee<h1>")
	})
	r.GET("/hello", func(c *gee.Context) {
		// expect /hello?name=XXX
		c.String(http.StatusOK, "Hello %s, you're at %s\n", c.GetQuery("name"), c.Path)
	})
	r.POST("/login", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}
