package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"gitee.com/bankarian/gee-web/gee"
	"gitee.com/bankarian/gee-web/gee/middlewares"
)

func main() {
	e := gee.New()
	e.Use(middlewares.Logger())
	e.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	e.LoadHTMLGlob("templates/*")
	e.Static("/assets", "./static")
	e.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "css.html", nil)
	})
	e.GET("/date", func(c *gee.Context) {
		c.HTML(http.StatusOK, "custom_func.html", gee.H{
			"title": "gee",
			"now":   time.Date(2021, 1, 10, 0, 0, 0, 0, time.UTC),
		})
	})
	e.Run(":9999")
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}
