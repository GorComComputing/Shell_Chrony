package main

import (
    "fmt"
    "os/exec"
    "os"
    "io"
    
    "bufio"
    "io/ioutil"
    "log"
    "strings"
)

var data = make([]byte, 64)


type Config struct {
	leapsectz	string
	driftfile    	string
	makestep       	string
	rtcsync 	bool
	logdir       	string
	local    	string
	server		string
	allow 		string
}

var config Config


func tst(words []string) {
	for idx, word := range words {
		fmt.Printf("Word %d is: %s\n", idx, word)
	}
}

func ls(words []string) {
	if len(words) < 2 {
		fmt.Println("Too little parameters")
	} else {
		cmd := exec.Command(words[0], words[1])
		out, err := cmd.Output()
		if err != nil {
			fmt.Println("could not run command: ", err)
		}
		fmt.Println(string(out))
	}
}

func start() {
	cmd := exec.Command("chronyd")
	//cmd := exec.Command("/etc/init.d/chrony", "start")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("start FAIL: ", err)
	} else {
		fmt.Println("start OK")
	}
	if len(out) > 0 {fmt.Println(string(out))}
}

func chrony_stop() {
	cmd := exec.Command("/etc/init.d/chrony", "stop")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("could not run command: ", err)
	}
	if len(out) > 0 {fmt.Println(string(out))}
}

func chrony_restart() {
	cmd := exec.Command("/etc/init.d/chrony", "restart")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("could not run command: ", err)
	}
	if len(out) > 0 {fmt.Println(string(out))}
}


func config_read(words []string) {
	file, err := os.Open("/etc/pzg-chrony.conf")
    	if err != nil{
        	fmt.Println(err) 
        	os.Exit(1) 
    	}
    	defer file.Close() 
     
    	
     
    	for{
        	n, err := file.Read(data)
        	if err == io.EOF{   	// если конец файла
            		break           // выходим из цикла
        	}
        	fmt.Print(string(data[:n]))
    	}
}

func config_write(words []string) {
	// собираем строку
	copy(words[0:], words[0+1:])
	words[len(words)-1] = ""
	words = words[:len(words)-1]
	text := strings.Join(words, " ")
	text = fmt.Sprintf("%s%s", text, "\n")
	
	cmd := exec.Command("cp", "/etc/pzg-chrony.conf", "/tmp/pzg-chrony.conf")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("could not copy: ", err)
	}
	fmt.Println(string(out))
    	  	
    	file, err := os.OpenFile("/tmp/pzg-chrony.conf", os.O_APPEND|os.O_WRONLY, 0600)
    	if err != nil {
        	fmt.Println("Unable to open file:", err) 
        	os.Exit(1) 
    	}
    	defer file.Close()

    	if _, err = file.WriteString(text); err != nil {
    		fmt.Println("Unable to write string:", err) 
        	os.Exit(1) 
    	} else {
    		cmd := exec.Command("cp", "/tmp/pzg-chrony.conf", "/etc/pzg-chrony.conf")
		out, err := cmd.Output()
		if err != nil {
			fmt.Println("could not back copy: ", err)
		}
		fmt.Println(string(out))
    		fmt.Println("Done.")
    	}
    	chrony_restart()
}

func replace(words []string) {
   stringNeeded := "server 192.168.1.1"
   stringToReplace := "server 192.168.3.3"
   filePath := "/tmp/config.txt"

   file, err := os.Open(filePath)
   if err != nil {
      log.Fatal(err)
   }
   defer file.Close()

   scanner := bufio.NewScanner(file)
   var lines []string
   for scanner.Scan() {
      text := scanner.Text()
      if scanner.Text() == stringNeeded {
         text = stringToReplace
      }

      lines = append(lines, text)
   }

   if err := scanner.Err(); err != nil {
      log.Fatal(err)
   }

   err = ioutil.WriteFile(filePath, []byte(strings.Join(lines, "\n")), 0644)
   if err != nil {
      log.Fatalln(err)
   }
   fmt.Println("Done.")
}

func cp(words []string) {
	cmd := exec.Command("cp", "config.txt", "/tmp/config.txt")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("could not run command: ", err)
	}
	fmt.Println("Output: ", string(out))
}

func bcp(words []string) {
	cmd := exec.Command("cp", "/tmp/config.txt", "config.txt")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("could not run command: ", err)
	}
	fmt.Println("Output: ", string(out))
}

func restore(words []string) {
	cmd := exec.Command("cp", "/root/pzg-chrony.conf", "/etc/pzg-chrony.conf")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("could not restore copy: ", err)
	}
	fmt.Println(string(out))
}


func scan() {
	// open the file
	file, err := os.Open("/etc/pzg-chrony.conf")

	//handle errors while opening
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}

	fileScanner := bufio.NewScanner(file)
	

	// read line by line
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if len(line) > 0 && string(line[0]) != "#" {
		
		words := strings.Fields(fileScanner.Text())
		
		switch words[0] {
   		case "leapsectz": config.leapsectz = words[1]
		case "driftfile": config.driftfile = words[1]
		case "makestep": config.makestep = words[1]
		case "rtcsync": config.rtcsync = true
		case "logdir": config.logdir = words[1]
		case "local": config.local = words[1]
		case "server": config.server = words[1]
		case "allow": config.allow = words[1]
		
		default: fmt.Println("Unknown directive")
		}
		
		
		//fmt.Println(words)
		//fmt.Println(fileScanner.Text())
		}
	}
	
	fmt.Println(config)
	
	// handle first encountered error while reading
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}

	file.Close()
}

func save() {
	
    	  	
    	file, err := os.OpenFile("./tmp.conf", os.O_WRONLY, 0600)
    	if err != nil {
        	fmt.Println("Unable to open file:", err) 
        	os.Exit(1) 
    	}
    	defer file.Close()

	
	if config.leapsectz != "" { 
    	if _, err = file.WriteString("leapsectz " + config.leapsectz + "\n"); err != nil {
    		fmt.Println("Unable to write string:", err) 
        	os.Exit(1) 
    	}
    	}
    	if config.driftfile != "" { 
    	if _, err = file.WriteString("driftfile " + config.driftfile + "\n"); err != nil {
    		fmt.Println("Unable to write string:", err) 
        	os.Exit(1) 
    	}
    	}
    	if config.makestep != "" { 
    	if _, err = file.WriteString("makestep " + config.makestep + "\n"); err != nil {
    		fmt.Println("Unable to write string:", err) 
        	os.Exit(1) 
    	}
    	}
    	if config.rtcsync { 
    	if _, err = file.WriteString("rtcsync\n"); err != nil {
    		fmt.Println("Unable to write string:", err) 
        	os.Exit(1) 
    	}
    	}
    	if config.logdir != "" { 
    	if _, err = file.WriteString("logdir " + config.logdir + "\n"); err != nil {
    		fmt.Println("Unable to write string:", err) 
        	os.Exit(1) 
    	}
    	}
    	if config.local != "" { 
    	if _, err = file.WriteString("local " + config.local + "\n"); err != nil {
    		fmt.Println("Unable to write string:", err) 
        	os.Exit(1) 
    	}
    	}
    	if config.server != "" { 
    	if _, err = file.WriteString("server " + config.server + "\n"); err != nil {
    		fmt.Println("Unable to write string:", err) 
        	os.Exit(1) 
    	}
    	}
    	if config.allow != "" { 
    	if _, err = file.WriteString("allow " + config.allow + "\n"); err != nil {
    		fmt.Println("Unable to write string:", err) 
        	os.Exit(1) 
    	}
    	}

}


