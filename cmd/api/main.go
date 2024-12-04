package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/gweningwarr/petOne/internal/app/api"
	"log"
)

var (
	configPath string = "configs/api.toml"
)

func init() {
	flag.StringVar(&configPath, "path", "configs/api.toml", "путь к конфигурации приложения")
}

func main() {
	log.Println("starting api")

	flag.Parse()

	config := api.NewConfig()
	if _, errTomlDecode := toml.DecodeFile(configPath, config); errTomlDecode != nil {
		log.Printf("Не удалось получить конфигурацию: %s", errTomlDecode)
	}

	server := api.New(config)

	log.Fatal(server.Start())

}
