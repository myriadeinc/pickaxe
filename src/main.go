package main

import (
	"os"
	"github.com/myriadeinc/pickaxe/src/api"
	"github.com/myriadeinc/pickaxe/src/util"
)

func initializeService() (bool) {
	//Initialize our Logger
	Logger.Init()
	MoneroApi.Init("http://0.0.0.0:8040")
	return true
}

func main() {

	var success bool = initializeService()
	if (success){
		// Early exit
		os.Exit(1)
	}
	// Starting PickAxe service
	Logger.Logger.Info("Starting PickAxe service")
}