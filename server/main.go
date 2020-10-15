package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

const dir = "./html"

func main() {
	fs := http.FileServer(http.Dir(dir))
	log.Print("Serving " + dir + " on http://localhost:8080")
	http.ListenAndServe(":8080", http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		fmt.Println(fmt.Sprintf("requesting %s", req.URL.Path))
		resp.Header().Add("Cache-Control", "no-cache")
		if strings.HasSuffix(req.URL.Path, ".wasm") {
			resp.Header().Set("content-type", "application/wasm")
		}
		if strings.HasSuffix(req.URL.Path, "index.js") {
			resp.Header().Set("content-type", "application/javascript")
		}
		fs.ServeHTTP(resp, req)
	}))
}
