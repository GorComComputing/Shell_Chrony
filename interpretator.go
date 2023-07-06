package main

import (
    //"fmt"
)


// Command list for interpretator
var cmd =  map[string]func([]string)string{
	"tst": cmd_tst,
	"ls": cmd_ls,
	"start": cmd_start,
        "stop": cmd_stop,
	"restart": cmd_restart,
	/*"read": config_read,
        "write": config_write,
        "replace": replace,
        "cp": cp,
        "bcp": bcp,
        "restore": restore,
        "scan": scan,
        "save": save,*/
	"quit": cmd_quit,
	"isactive": cmd_isActive,
	"activity": cmd_activity,
	"tracking": cmd_tracking,
	"sources": cmd_sources,
	"sourcestats": cmd_sourcestats,
	"clients": cmd_clients,
	"config": cmd_config,
	"restore": cmd_restore,
	"write": cmd_write,
	"save": cmd_save,
	

}



// Interpretator 
func interpretator(words []string) string {
	if _, ok := cmd[words[0]]; ok {
		return cmd[words[0]](words)
	} else{
		//fmt.Println("Unknown command: " + words[0])
		return "Unknown command: " + words[0] + "\n"
	}
}


