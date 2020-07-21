package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/elazarl/goproxy"
)

func main() {
	addr := flag.String("addr", ":8080", "proxy listen address")
	flag.Parse()

	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	proxy.NonproxyHandler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path == "/ping" {
			fmt.Fprintf(w, "pong")
		} else {
			req.URL.Host = req.Host
			proxy.ServeHTTP(w, req)
		}
	})
	if err := http.ListenAndServe(*addr, proxy); err != nil {
		panic(err)
	}
}
