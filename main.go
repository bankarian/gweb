package main

import (
	"net/http"

	"gitee.com/bankarian/gee-web/gee"
)

func main() {

	r := gee.NewEngine()
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page<h1>")
	})
	r.POST("/login", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *gee.Context) {
			c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
		})
		v1.GET("/hello", func(c *gee.Context) {
			// expect /hello?name=XXX
			c.String(http.StatusOK, "Hello %s, you're at %s\n", c.GetQuery("name"), c.Path)
		})
	}
	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *gee.Context) {
			c.String(http.StatusOK, "Hello %s, you're at %s\n", c.Param("name"), c.Req.URL.Path)
		})

		v2.GET("/assets/*filepath", func(c *gee.Context) {
			c.JSON(http.StatusOK, gee.H{
				"filepath": c.Param("filepath"),
			})
		})
	}

	r.Run(":9999")
}
