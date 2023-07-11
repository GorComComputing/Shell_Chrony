package main

import (
	"fmt"
	"net/http"
)

// Errors handler to STDIO
func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		panic(err)
	}
}

// Errors handler to HTTP
func CheckErrorHttp(err error, w http.ResponseWriter) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
