package tor

import (
	"net"
	"fmt"
	"os"
	"context"
	"path/filepath"
	bine "github.com/cretz/bine/tor"
)

func TorClient(onionUrl string)(t *bine.Tor, conn net.Conn){

	ctx := context.Background()

	t, err := bine.Start(ctx, &bine.StartConf{
		ExePath: filepath.FromSlash(ReadAppConfigFile().App.TorPath),
		DebugWriter: os.Stdout,
		TempDataDirBase: filepath.FromSlash(ReadAppConfigFile().App.AppSessionsDir),
		// DataDir: filepath.FromSlash(ReadAppConfigFile().App.AppSessionsDir),
	})
	if err != nil{
		fmt.Println("Error starting Tor:", err)
		return
	}
	
	err = t.EnableNetwork(ctx, true)
	if err != nil{
		fmt.Println("Error enabling network:", err)
		return
	}

	dialer, err := t.Dialer(ctx, nil)
	if err != nil{
		fmt.Println("Error creating Tor dialer", err)
		return
	}

	conn, err = dialer.DialContext(ctx, "tcp", onionUrl + ":80")
	if err != nil{
		fmt.Println("Error at onion dialer", err)
		return
	}

	fmt.Println("Connected to:", onionUrl)

	return t, conn
}