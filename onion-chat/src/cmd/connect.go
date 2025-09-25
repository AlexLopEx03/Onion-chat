package commands

import (
	"fmt"
	"onion-chat/src/tor"
	"onion-chat/src/chat"
)

func Connect(args [] string){

	t, conn:= tor.TorClient(args[2])
	if t == nil || conn == nil{
		fmt.Println("Error, couldn't connect to", args[2])
		return
	}
	defer func(){
		_ = conn.Close()
		_ = t.Close()
	}()

	chat.ClientMessage(conn)
}