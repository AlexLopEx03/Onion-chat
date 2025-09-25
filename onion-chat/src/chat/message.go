package chat

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"github.com/common-nighthawk/go-figure"
)

func ServerMessage(conn net.Conn){
	go func(){
		fmt.Print("\033[2J")
		fmt.Print("\033[H")
		figure.NewFigure("Onion chat", "", true).Print()
		fmt.Println()
		reader := bufio.NewReader(conn)
		for{
			message, err := reader.ReadString('\n')
			if err != nil{
				fmt.Println("Client disconnected:", err)
				os.Exit(0)
			}
			fmt.Print("\033[32m[Anonymous] " + message + "\033[0m")
			// _, _ = conn.Write([] byte ("Message received\n"))
		}
	}()

	for{
		text, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil{
			fmt.Println("Error reading stdin:", err)
			return
		}
		
		_, err = conn.Write([] byte (text))
		if err != nil{
			fmt.Println("Error sending message:", err)
			return
		}
	}

	// Add defer ending message
}

func ClientMessage(conn net.Conn){
	go func() {
		fmt.Print("\033[2J")
		fmt.Print("\033[H")
		figure.NewFigure("Onion chat", "", true).Print()
		fmt.Println()
		reader := bufio.NewReader(conn)
		for {
			message, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Server disconnected:", err)
				os.Exit(0)
			}
			fmt.Print("\033[32m[Anonymous] " + message + "\033[0m")
			// _, _ = conn.Write([] byte ("Message received\n"))
		}
	}()

	for{
		text, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil{
			// fmt.Println("Error reading stdin:", err)
			fmt.Println("User disconnected:", err)
			return
		}
		_, err = conn.Write([] byte (text))
		if err != nil{
			fmt.Println("Error sending message:", err)
			return
		}
	}

	// Add defer ending message
}