package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "net/http"
)


var exit_status bool = true	// false = exit


func main() {
	var cmd_line string
	var words = make([]string, len(os.Args)-1)

	// pars command line args
	if len(os.Args) > 1 {

		// TUI mode
		if os.Args[1] == "-w"{
                        fmt.Println("Window")
                        os.Exit(0)
                }

        	copy(words[0:], os.Args[1:])
                exit_status = false
                
        	out := interpretator(words)
		if len(out) > 0 {
			fmt.Print(out)
		}
        	
		os.Exit(0)
	}

	// start web-server
	http.HandleFunc("/api", http_pars)
	fmt.Println("WebServer started OK. Try http://192.168.1.136:8084")
	go http.ListenAndServe(":8084", nil)

	// start shell
	for exit_status {
		fmt.Print("Chrony Shell> ")
		// ввод строки с пробелами
    		scanner := bufio.NewScanner(os.Stdin)
    		scanner.Scan()
    		cmd_line = scanner.Text()
    		// разбиение на подстроки по пробелу
    		words = strings.Fields(cmd_line)

		out := interpretator(words)
		if len(out) > 0 {
			fmt.Print(out)
		}
	}
}


func cmd_quit(words []string) string {
	exit_status = false
	return ""
}
