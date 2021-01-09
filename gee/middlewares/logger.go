package middlewares

import (
	"log"
	"time"

	"gitee.com/bankarian/gee-web/gee"
)


func Logger() gee.HandlerFunc {
	return func(c *gee.Context) {
		t := time.Now()
		// Process request
		c.Next()
		// Calculate resolution time
		log.Printf("[%d] %s in %v\n", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}