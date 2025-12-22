package main

import (
	"fmt"

	"github.com/tengenatari/web-referee/config"
	"github.com/tengenatari/web-referee/internal/bootstrap"
)

func main() {
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		panic(fmt.Sprintf("Error loading config: %v", err))
	}
	storage := bootstrap.InitPGStorage(cfg)
	producer := bootstrap.InitMessageProducer(cfg)
	service := bootstrap.InitWebRefereeService(storage, producer)
	api := bootstrap.InitWebRefereeServiceAPI(service)

	bootstrap.AppRun(api, cfg)

}
