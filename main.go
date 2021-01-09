package main

import (
	"log"
	"net/http"
	"time"

	"gitee.com/bankarian/gee-web/gee"
	"gitee.com/bankarian/gee-web/gee/middlewares"
)

func main() {

	engine := gee.New()
	engine.Use(middlewares.Logger()) // global middleware Logger
	engine.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page<h1>")
	})
	engine.POST("/login", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	v2 := engine.Group("/v2")
	v2.Use(func(c *gee.Context) {
		t := time.Now()
		c.Fail(http.StatusInternalServerError, "Internal Server Error")
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	})
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

	engine.Run(":9999")
}
