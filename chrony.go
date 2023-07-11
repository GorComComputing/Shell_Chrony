package main

import (
    "fmt"
    "os/exec"
    "os"
    //"io"
    "bufio"
    //"io/ioutil"
    "log"
    "strings"
)

var data = make([]byte, 64)

// Структура config-файла
type Config struct {
	Leapsectz	string
	Driftfile    	string
	Makestep       	string
	Makestep2      	string
	Rtcsync 	bool
	Logdir       	string
	Local    	string
	Server		string
	Allow 		string
}


var config Config


func cmd_tst(words []string) string {
	for idx, word := range words {
		fmt.Printf("Word %d is: %s\n", idx, word)
	}
	return ""
}

func cmd_ls(words []string) string {
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
	return ""
}


// Запуск Chrony
func cmd_start(words []string) string {
	var output string
	cmd := exec.Command("chronyd")
	//cmd := exec.Command("/etc/init.d/chrony", "start")
	out, err := cmd.Output()
	if err != nil {
		output = "start FAIL: " + err.Error() + "\n"
	} else {
		output = "start OK\n"
	}
	if len(out) > 0 {fmt.Println(string(out))}
	return output
}


// Остановка Chrony
func cmd_stop(words []string) string {
	var output string
	cmd := exec.Command("killall", "chronyd")
	out, err := cmd.Output()
	if err != nil {
                output = "stop FAIL: " + err.Error() + "\n"
        } else {
                output = "stop OK\n"
        }
	if len(out) > 0 {fmt.Println(string(out))}
	return output
}


// Перезапуск Chrony
func cmd_restart(words []string) string {
	output1 := cmd_stop(words)
	output2 := cmd_start(words)
	return output1 + output2
}


// Проверяет запущен ли Chrony
func isActive() bool {
        file, err := os.Open("/var/run/chrony/chronyd.pid")
        defer file.Close()

        //handle errors while opening
        if err != nil {
                chan_isActive <- string("no")
                return false
        } else {
                chan_isActive <- string("yes")
                return true
        }
}


// Проверяет запущен ли Chrony
func cmd_isActive(words []string) string {
	if isActive() {
		return "yes\n"
	} else {
		return "no\n"
	}
}


// Читает config-файл
func scan() (Config, string) {

	var config Config

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
   			case "leapsectz": config.Leapsectz = words[1]
			case "driftfile": config.Driftfile = words[1]
			case "makestep": {config.Makestep = words[1]; config.Makestep2 = words[2]}
			case "rtcsync": config.Rtcsync = true
			case "logdir": config.Logdir = words[1]
			case "local": config.Local = words[2]
			case "server": config.Server = words[1]
			case "allow": config.Allow = words[1]
		
			default: fmt.Println("Unknown directive")
			}
		}
	}
	
	// handle first encountered error while reading
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}

	defer file.Close()
	
	dat, err := os.ReadFile("/etc/pzg-chrony.conf")
	return config, string(dat)
}


// Проверяет активность Chrony
func cmd_activity(words []string) string {
	cmd := exec.Command("chronyc", "activity")
	out, _ := cmd.Output()
	/*if err != nil {
		fmt.Println("could not run command: ", err)
	}*/
	return string(out)
}


// chronyc tracking - информация о сервере времени
func cmd_tracking(words []string) string {
	cmd := exec.Command("chronyc", "tracking")
	out, _ := cmd.Output()
	/*if err != nil {
		fmt.Println("could not run command: ", err)
	}*/
	return string(out)
}


// chronyc sources -v   список источников времени
func cmd_sources(words []string) string {
	cmd := exec.Command("chronyc", "sources", "-v")
	out, _ := cmd.Output()
	/*if err != nil {
		fmt.Println("could not run command: ", err)
	}*/
	return string(out)
}


// chronyc sourcestats -v список источников времени
func cmd_sourcestats(words []string) string {
	cmd := exec.Command("chronyc", "sourcestats", "-v")
	out, _ := cmd.Output()
	/*if err != nil {
		fmt.Println("could not run command: ", err)
	}*/
	return string(out)
}


// chronyc clients - список подключенных клиентов
func cmd_clients(words []string) string {
	cmd := exec.Command("chronyc", "clients")
	out, _ := cmd.Output()
	/*if err != nil {
		fmt.Println("could not run command: ", err)
	}*/
	return string(out)
}


// Читает Config-файл
func cmd_config(words []string) string {
    	_ , File := scan()
    	File = fmt.Sprintf("%s%s", File, "\n")
    	return File
    	//messages <- string("Config-файл прочитан")
}


// Восстановление config-файла
func cmd_restore(words []string) string {
    	// перенести из files в основной файл
    	cmd := exec.Command("cp", "./files/pzg-chrony.conf", "/etc/pzg-chrony.conf")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("could not back copy: ", err)
	}
	if len(out) > 0 {fmt.Println(string(out))}
    	//fmt.Println("Restored OK")
    	cmd_restart(words)
    	
    	return "Config restored OK\n"
    	//_ , File := scan()
    	//fmt.Fprintf(w, File)
    	//messages <- string("Config-файл Chrony восстановлен<br/>Chrony запущен")
}


// Сохраняет config-файл и перезапускает Chrony
func cmd_write(words []string) string {
	var output string
	params := make(map[string]string)
	for i := 1; i < len(words); i++ { 
		params[strings.SplitAfter(words[i], "=")[0]] = strings.SplitAfter(words[i], "=")[1]
	}
	
	// parameters from POST
	Leapsectz := params["leapsectz="]
	Driftfile := params["driftfile="]
	Makestep := params["makestep="]
	Makestep2 := params["makestep2="]
	Rtcsync := params["rtcsync="]
	//fmt.Println(Rtcsync)
	Logdir := params["logdir="]
	Local := params["localStratum="]
	Server := params["server="]
	Allow := params["allow="]
	
    	file, err := os.OpenFile("./files/tmp.conf", os.O_TRUNC | os.O_WRONLY, 0600)
    	if err != nil {
        	//fmt.Println("Unable to open file:", err) 
        	output = "Unable to open file: " + err.Error() + "\n"
        	return output
        	//os.Exit(1) 
    	}
    	defer file.Close()

	if Leapsectz != "" { 
    	if _, err = file.WriteString("leapsectz " + Leapsectz + "\n"); err != nil {
    		//fmt.Println("Unable to write string:", err) 
        	output = "Unable to write string: " + err.Error() + "\n"
        	return output
        	//os.Exit(1) 
    	}
    	}
    	if Driftfile != "" { 
    	if _, err = file.WriteString("driftfile " + Driftfile + "\n"); err != nil {
    		//fmt.Println("Unable to write string:", err) 
        	output = "Unable to write string: " + err.Error() + "\n"
        	return output
        	//os.Exit(1) 
    	}
    	}
    	if Makestep != "" && Makestep2 != "" { 
    	if _, err = file.WriteString("makestep " + Makestep + " " + Makestep2 + "\n"); err != nil {
    		//fmt.Println("Unable to write string:", err) 
        	output = "Unable to write string: " + err.Error() + "\n"
        	return output
        	//os.Exit(1) 
    	}
    	}
    	if Rtcsync != "" { 
    	if _, err = file.WriteString("rtcsync\n"); err != nil {
    		//fmt.Println("Unable to write string:", err) 
        	output = "Unable to write string: " + err.Error() + "\n"
        	return output
        	//os.Exit(1) 
    	}
    	}
    	if Logdir != "" { 
    	if _, err = file.WriteString("logdir " + Logdir + "\n"); err != nil {
    		//fmt.Println("Unable to write string:", err) 
        	output = "Unable to write string: " + err.Error() + "\n"
        	return output
        	//os.Exit(1) 
    	}
    	}
    	if Local != "" { 
    	if _, err = file.WriteString("local stratum " + Local + "\n"); err != nil {
    		//fmt.Println("Unable to write string:", err) 
        	output = "Unable to write string: " + err.Error() + "\n"
        	return output
        	//os.Exit(1)  
    	}
    	}
    	if Server != "" { 
    	if _, err = file.WriteString("server " + Server + "\n"); err != nil {
    		//fmt.Println("Unable to write string:", err) 
        	output = "Unable to write string: " + err.Error() + "\n"
        	return output
        	//os.Exit(1) 
    	}
    	}
    	if Allow != "" { 
    	if _, err = file.WriteString("allow " + Allow + "\n"); err != nil {
    		//fmt.Println("Unable to write string:", err) 
        	output = "Unable to write string: " + err.Error() + "\n"
        	return output
        	//os.Exit(1) 
    	}
    	}
    	
    	// перенести из tmp в основной файл
    	cmd := exec.Command("cp", "./files/tmp.conf", "/etc/pzg-chrony.conf")
	out, err := cmd.Output()
	if err != nil {
		//fmt.Println("could not back copy: ", err)
		output = "Could not back copy: " + err.Error() + "\n"
        	return output
	}
	if len(out) > 0 {fmt.Println(string(out))}
    	//fmt.Println("Saved OK")
    	cmd_restart(words)
    	
    	
    	_ , File := scan()
	//File = strings.Replace(File, "\n", "<br/>", -1)
	//bks.File = template.HTML(File)
    	
    	//fmt.Fprintf(w, File)
    	//messages <- string("Config-файл Chrony сохранен<br/>Chrony запущен")
    	File = fmt.Sprintf("%s%s", File, "\n")
    	return File
}


// Сохраняет Config-файл
func cmd_save(words []string) string {
	var output string
	text := strings.SplitAfter(words[1], "=")[1]
				
    	file, err := os.OpenFile("./files/tmp.conf", os.O_TRUNC | os.O_WRONLY, 0600)
    	if err != nil {
        	output = "Unable to open file: " + err.Error() + "\n"
        	return output
    	}
    	defer file.Close()

	if text != "" { 
    	if _, err = file.WriteString(text); err != nil {
    		output = "Unable to write string: " + err.Error() + "\n"
    		return output
    	}
    	}
    	
    	// перенести из tmp в основной файл
    	cmd := exec.Command("cp", "./files/tmp.conf", "/etc/pzg-chrony.conf")
	_, err = cmd.Output()
	if err != nil {
		output = "Could not back copy: " + err.Error() + "\n"
    		return output
	}
    	//fmt.Println("Saved OK")
    	cmd_restart(words)
    	
    	_ , File := scan()
    	
    	//fmt.Fprintf(w, File)
    	//fmt.Println(File)
    	//messages <- string("Config-файл Chrony сохранен<br/>Chrony запущен")
    	File = fmt.Sprintf("%s%s", File, "\n")
    	return File
}

