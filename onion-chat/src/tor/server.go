package tor

import (
	"fmt"
	"context"
	"onion-chat/src/chat"

	bine "github.com/cretz/bine/tor"
)

func TorServer(onionService *bine.OnionService){
	for {
		conn, err := onionService.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		fmt.Println("\033[32m New client connected\033[0m")

		go chat.ServerMessage(conn)
	}
}

func StartTor(t *bine.Tor, torConfig *bine.ListenConf) (*bine.OnionService, string){

	listener, err := t.Listen(context.Background() ,torConfig)
	if err != nil{
		panic(err)	
	}
	
	return listener, listener.ID + ".onion"
}