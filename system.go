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
		output = "top FAIL: " + err.Error() + "\n"
	} else {
		output = string(out)
	}
	return output 
}

// Монитор ресурсов и процессов (top)
func cmd_netstat(words []string) string {
	var output string
	cmd := exec.Command("netstat", "-a")
	out, err := cmd.Output()
	if err != nil {
		output = "netstat FAIL: " + err.Error() + "\n"
	} else {
		output = string(out)
	}
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



