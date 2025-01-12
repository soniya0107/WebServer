package main

import (
	"fmt"
	"net/http"
)

func formHandler(resp http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/form" {
		http.Error(resp, "404 not found", http.StatusNotFound)
		return
	}

	if req.Method != "POST" {
		http.Error(resp, "Method not supported", http.StatusNotFound)
		return
	}

	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(resp, "Error while parsing form: %v", err)
		return
	}

	fmt.Println("POST Successful")
	name := req.FormValue("name")
	address := req.FormValue("address")

	fmt.Fprintf(resp,"YAY\nName: %v\nAddress: %v", name, address)
	fmt.Printf("Name: %v\nAddress: %v", name, address)
}
