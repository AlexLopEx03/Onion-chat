package main

import (
	"fmt"
	"os"
	"onion-chat/src/cmd"
)

func main(){

	if len(os.Args) < 2{
		println("No command provided.\nRun --help to get commands info.")
		return
	}

	switch os.Args[1]{
	case "connect":
		commands.Connect(os.Args)
	case "create":
		commands.Create(os.Args)		
	case "help", "-h", "--help":
		commands.Help(os.Args)
	case "uninstall":
		commands.Uninstall(os.Args)
	// case "update":
	// 	commands.UpdateTor(os.Args)
	default:
		fmt.Println("Error, unknown command.\nRun --help to check syntax.")
	}
}