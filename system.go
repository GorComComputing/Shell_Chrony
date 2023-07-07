package main

import (
    //"fmt"
    "os/exec"
    //"strings"
    "net/http"
    "io/ioutil"
)


// Выполнить команду
func cmd_run(words []string) string {
	run_cmd := words[1]
	if len(words) > 2 {
		copy(words[0:], words[2:])
		words = words[:len(words)-2]
		
		cmd := exec.Command(run_cmd, words...)
		out, _ := cmd.Output()
		/*if err != nil {
			fmt.Println("could not run command: ", err)
		}*/
		//fmt.Println(string(out))
		return string(out)
	} else {
		cmd := exec.Command(run_cmd)
		out, _ := cmd.Output()
		/*if err != nil {
			fmt.Println("could not run command: ", err)
		}*/
		//fmt.Println(string(out))
		return string(out)
	}
}



// Монитор ресурсов и процессов (top)
func cmd_top(words []string) string {
	var output string
	cmd := exec.Command("top", "-n", "1")
	out, err := cmd.Output()
	if err != nil {
		//fmt.Println("start FAIL: ", err)
		output = "top FAIL: " + err.Error() + "\n"
	} else {
		//fmt.Println("start OK")
		output = string(out)
	}
	//if len(out) > 0 {fmt.Println(string(out))}
	return output 
}

// Монитор ресурсов и процессов (top)
func cmd_netstat(words []string) string {
	var output string
	cmd := exec.Command("netstat", "-a")
	out, err := cmd.Output()
	if err != nil {
		//fmt.Println("start FAIL: ", err)
		output = "netstat FAIL: " + err.Error() + "\n"
	} else {
		//fmt.Println("start OK")
		output = string(out)
		//if len(out) > 0 {fmt.Println(string(out))}
	}
	//if len(out) > 0 {fmt.Println(string(out))}
	return output 
}



// Завершить процесс по PID
func cmd_kill(words []string) string {
	var output string
	pid_name := words[1]
	cmd := exec.Command("kill", "-9", pid_name)
	_, err := cmd.Output()
	if err != nil {
		output = "kill " + pid_name + " FAIL: " + err.Error() + "\n"
	} else {
		output = pid_name + " killed OK\n"
	}
	return output 
}


// Завершить процесс по Name
func cmd_killall(words []string) string {
	var output string
	pid_name := words[1]
	cmd := exec.Command("killall", pid_name)
	_, err := cmd.Output()
	if err != nil {
		output = "killall " + pid_name + " FAIL: " + err.Error() + "\n"
	} else {
		output = pid_name + " killed OK\n"
	}
	return output 
}


func MakeRequest(words []string) string {
	var output string
	
	db_host := "http://192.168.1.136"
	db_port := "8086"
	
	db_Query := "select"
	
	resp, err := http.Get(db_host + ":" + db_port + "/api?cmd=" + db_Query)
	if err != nil {
		output = "Request FAIL: " + err.Error() + "\n"
		return output
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		output = "Request FAIL: " + err.Error() + "\n"
		return output
	}

	//reqArr := strings.Fields(string(body))
	//fmt.Println(reqArr)

	return string(body) 
}


