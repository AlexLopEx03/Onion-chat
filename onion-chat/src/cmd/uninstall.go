package commands

import (
	"fmt"
	"os"
	"os/exec"
)

func Uninstall(args [] string){
	command := exec.Command(
		"powershell", "-ExecutionPolicy", "Bypass", "-Command",
		`Start-Process powershell -ArgumentList '-ExecutionPolicy Bypass -Command "iwr https://raw.githubusercontent.com/alexlopex03/Onion-chat/main/scripts/uninstaller.ps1 | iex"' -WindowStyle Hidden`,
	)

	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	command.Stdin = os.Stdin

	err := command.Run()
	if err != nil{
		fmt.Println("Error while uninstalling the App")
	}
}	