package commands

import (
	"fmt"
	"github.com/common-nighthawk/go-figure"
)

func Help(args [] string){
	
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
	figure.NewFigure("Onion chat", "", true).Print()

	fmt.Println()

	fmt.Print("     Flags             ")
	fmt.Print("Arguments     ")
	fmt.Print("Description \n")

	fmt.Println()

	fmt.Print("         ")
	fmt.Print("--create      ")
	fmt.Print("              ")
	fmt.Print("Creates an onion-chat service and returns the onion url\n")

	fmt.Print("         ")
	fmt.Print("--connect     ")
	fmt.Print("[onion-url]   ")
	fmt.Print("Connects to an existing onion-chat service with an onion url\n")

	fmt.Print("         ")
	fmt.Print("--update      ")
	fmt.Print("              ")
	fmt.Print("Updates tor engine and onion-chat.exe\n")

	fmt.Print("         ")
	fmt.Print("--uninstall   ")
	fmt.Print("              ")
	fmt.Print("Uninstalls onion-chat app\n")

	fmt.Print("-h       ")
	fmt.Print("--help        ")
	fmt.Print("              ")
	fmt.Print("Shows all available commands\n")
}