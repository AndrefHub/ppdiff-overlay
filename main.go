package main

import (
	"flag"

	"github.com/AndrefHub/ppdiff-overlay/token"
	"github.com/AndrefHub/ppdiff-overlay/web"
)

func main() {
	// config.Init()
	// shouldWeUpdate := flag.Bool("autoupdate", true, "Should we auto update the application?")

	flag.Parse()

	// if *shouldWeUpdate {
	// 	updater.DoSelfUpdate()
	// }
	go token.SetUp()
	go web.SetupRoutes()

	web.HTTPServer()
}
