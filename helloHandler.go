package main

import (
	"fmt"
	"net/http"
)

func helloHandler(resp http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/hello" {
		http.Error(resp, "404 Not found", http.StatusNotFound)
	}

	if req.Method != "GET" {
		http.Error(resp, "Method not supported", http.StatusNotFound)
	}

	fmt.Fprintf(resp, "hello and welcome to our website !")
}
