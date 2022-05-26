package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"log"
	"server/internal/app/server"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/server.toml", "Path to config file")
}

func main() {

	flag.Parse()
	config := server.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	//start app
	serverGRPC := server.New(config)
	if err := serverGRPC.Start(); err != nil {
		log.Fatal(err)
	}
}
