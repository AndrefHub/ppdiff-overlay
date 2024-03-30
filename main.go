package main

import (
	"log"
	"os"

	"github.com/AndrefHub/ppdiff-overlay/config"
	"github.com/AndrefHub/ppdiff-overlay/token"
	"github.com/AndrefHub/ppdiff-overlay/updater"
	"github.com/AndrefHub/ppdiff-overlay/web"
	"github.com/spf13/cast"
)

func ChangeLogDestinationToFile() {
	f, err := os.OpenFile("ppdiff-overlay.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0664)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	log.SetOutput(f)
}

func main() {
	ChangeLogDestinationToFile()
	config.Init()

	if cast.ToBool(config.Config["SelfUpdate"]) {
		updater.DoSelfUpdate()
	}
	token.SetUp(config.Config["ClientID"], config.Config["ClientSecret"])
	web.HTTPServer()
}
