package main

import (
	"fmt"
	"github.com/ybbus/jsonrpc"
	"github.com/myriadeinc/pickaxe/src/api"
	"github.com/myriadeinc/pickaxe/src/util"
)

func initializeService() (int) {
	//Initialize our Logger
	Logger.Init()
	return 1
}

func main() {

	var success int = initializeService()
	if (0 == success){
		// Early exit
	}
	// Starting PickAxe service
	Logger.Logger.Info("Starting PickAxe service")

	var rpcClient jsonrpc.RPCClient = MoneroAPI.Init("http://0.0.0.0:8040")
	fmt.Println(rpcClient)
}