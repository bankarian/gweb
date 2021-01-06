package main

import (
	"fmt"
	"net/http"

	"gitee.com/gee-web/gee"
)

func main() {
	r := gee.New()
	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	})
	r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})

	r.Run(":9999")
}

// func main() {
// 	engine := new(Engine)
// 	log.Fatal(http.ListenAndServe(":9999", engine))
// }

// type Engine struct{}

// func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
// 	switch req.URL.Path {
// 		case "/":
// 			fmt.Fprintf(w, "Hello, URL.Path = %q\n", req.URL.Path)
// 		case "/hello":
// 			for k, v := range req.Header {
// 				fmt.Fprintf(w, "Headr[%q] = %q\n", k, v)
// 			}
// 		default:
// 			fmt.Fprintf(w, "404 Not Found: %s\n", req.URL)
// 	}
// }

// func main() {
// 	http.HandleFunc("/", indexHandler)
// 	http.HandleFunc("/hello", helloHandler)
// 	log.Fatal(http.ListenAndServe(":9999", nil))
// }

// func indexHandler(w http.ResponseWriter, req *http.Request) {
// 	fmt.Fprintf(w, "Hello, URL.Path = %q\n", req.URL.Path)
// }

// func helloHandler(w http.ResponseWriter, req *http.Request) {
// 	for k, v := range req.Header {
// 		fmt.Fprintf(w, "Headr[%q] = %q\n", k, v)
// 	}
// }
