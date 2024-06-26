package config

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/l3lackShark/config"
)

// Config file
var Config map[string]string

// Init the config file
func Init() {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)

	cfg, err := config.SetFile(filepath.Join(exPath, "config.ini"))
	if err == config.ErrDoesNotExist {
		d := []byte(`[Client]
ClientID = dummyID
ClientSecret = dummySecret
		
[Web]
serverip = 127.0.0.1:57275
cors = false

[Settings]
AutoUpdate = true

`)
		if err := ioutil.WriteFile(filepath.Join(exPath, "config.ini"), d, 0644); err != nil {
			panic(err)
		}
		cfg, err = config.SetFile(filepath.Join(exPath, "config.ini"))
		if err != nil {
			panic(err)
		}
	} else if err != nil {
		log.Fatalln(err)
	}
	Config, err = cfg.Parse()
	if err != nil {
		panic(err)
	}
}
