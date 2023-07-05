package main

import (
    "fmt"
    "os"
    
    "bufio"
    "strings"
    "net/http"
)

var exit_status bool = true


// http handler
func http_api(w http.ResponseWriter, r *http.Request) {
	// parameters from POST or GET
        r.ParseForm()
	words := []string{}

	for _, values := range r.Form {   // range over map
  		for _, value := range values {    // range over []string
     			words = append(words, value)
  		}
	}
	interpretator(words)
}


func main() {
	var cmd_line string
	var words = make([]string, len(os.Args)-1)

	if len(os.Args)>1 {
        	copy(words[0:], os.Args[1:])
                exit_status = false
        	interpretator(words)
		os.Exit(0)
	}

	// Page routs
	http.HandleFunc("/api", http_api)
	fmt.Println("WebServer started OK. Try http://192.168.1.136:8084")
	go http.ListenAndServe(":8084", nil)

	for exit_status {
		fmt.Print("Chrony Shell> ")
		// ввод строки с пробелами
    		scanner := bufio.NewScanner(os.Stdin)
    		scanner.Scan()
    		cmd_line = scanner.Text()
    		// разбиение на подстроки по пробелу
    		words = strings.Fields(cmd_line)
   		
		interpretator(words)
	}
}


// Interpretator 
func interpretator(words []string) {
		switch words[0] {
                case "tst": tst(words)
                case "ls": ls(words)
                case "start": start()
                case "stop": chrony_stop()
                case "restart": chrony_restart()
                case "read": config_read(words)
                case "write": config_write(words)
                case "replace": replace(words)
                case "cp": cp(words)
                case "bcp": bcp(words)
                case "restore": restore(words)
                case "scan": scan()
                case "save": save()

                case "quit": exit_status = false
		default: fmt.Println("Unknown command: " + words[0])
                }
}


