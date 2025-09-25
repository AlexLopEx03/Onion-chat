package commands

import (
	"fmt"
	"onion-chat/src/tor"
)

func Create(args [] string){

	t, torConfig := tor.TorConfig(tor.ReadAppConfigFile())
	onionService, onionUrl := tor.StartTor(t, torConfig)
	fmt.Println("\033[32m Tor server at:\033[0m", onionUrl)
	fmt.Println("\033[32m Waiting the client to connect...\033[0m")

	tor.TorServer(onionService)
}