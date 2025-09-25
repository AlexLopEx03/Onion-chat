package tor

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	bine "github.com/cretz/bine/tor"
	"gopkg.in/yaml.v3"
)

type AppConfig struct{
	App struct{
		Platform string `yaml:"platform"`
		AppDir string `yaml:"appDir"`
		AppSessionsDir string `yaml:"appSessionsDir"`
		AppPath string `yaml:"appPath"`
		TorPath string `yaml:"torPath"`
	} `yaml:"app"`
}

func ReadAppConfigFile()(*AppConfig){
	exePath, err := os.Executable()
	if err != nil{
		fmt.Println("Error while reading current .exe path")
		os.Exit(0)
	}
	exeDir:= filepath.Dir(exePath)
	configPath := filepath.Join(exeDir, "config.yaml")

	file, err := os.Open(configPath)
	if err != nil{
		fmt.Println("Error while reading config.yaml")
		os.Exit(0)
	}
	defer file.Close()

	var config AppConfig
	err = yaml.NewDecoder(file).Decode(&config)
	if err != nil{
		fmt.Println("Error while reading config.yaml")
		os.Exit(0)
	}

	return &config
}

func TorConfig(config *AppConfig)(*bine.Tor, *bine.ListenConf){

	t, err := bine.Start(context.Background(), &bine.StartConf{
		ExePath: filepath.FromSlash(ReadAppConfigFile().App.TorPath),
		DebugWriter: os.Stdout,
		TempDataDirBase: filepath.FromSlash(ReadAppConfigFile().App.AppSessionsDir),
		// DataDir: filepath.FromSlash(ReadAppConfigFile().App.AppSessionsDir),
	})
	if err != nil{
		panic(err)
	}

	torConfig := &bine.ListenConf{
		RemotePorts: [] int {80},
		LocalPort: 3000,
		Version3: true,
	}
	
	return t, torConfig
}