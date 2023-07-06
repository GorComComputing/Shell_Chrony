package main

import (
	"fmt"
    	"net/http"
	"strings"
 	"io"
)


// создание каналов для long poll
var chan_message chan string = make(chan string, 100)
var chan_isActive chan string = make(chan string, 100)


// /api handler
func http_pars(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)	// enable CORS
	// parameters from POST or GET
        r.ParseForm()
	words := []string{}

	for param, values := range r.Form {   // range over map
  		for _, value := range values {    // range over []string
     			if param == "cmd" {
				words = strings.Fields(value)
			} else {
				words = append(words, string(param) + "=" + string(value))
			}
  		}
	}
	//fmt.Println(words)
	out := interpretator(words)
	if len(out) > 0 {
		//fmt.Print(out)
		fmt.Fprintf(w, out)
	}
	

}


// Enable CORS
func enableCors(w *http.ResponseWriter) {
        (*w).Header().Set("Access-Control-Allow-Origin", "*")
}


// получение longpoll и установка канала для ответа
func PollMessage(w http.ResponseWriter, req *http.Request) {
    	io.WriteString(w, <-chan_message)
}


// получение longpoll и установка канала для ответа
func Poll_isActive(w http.ResponseWriter, req *http.Request) {
    	io.WriteString(w, <-chan_isActive)
}
