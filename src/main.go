package main

import (
	"os"
	"net/http"
	"fmt"
	"log"
	"github.com/gorilla/mux"
	"github.com/myriadeinc/pickaxe/src/util/logger"
	"github.com/myriadeinc/pickaxe/src/util/config"
	"github.com/myriadeinc/pickaxe/src/routes/subscriber"
	"github.com/myriadeinc/pickaxe/src/services/polling"
)

func healthcheck(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "OK")
}

func initializeService() (bool) {
	//Initialize our Logger
	LoggerUtil.Init()
	ConfigUtil.Init()
	PollingService.Init()
	return true
}

func main() {
	var success bool = initializeService()
	if (!success){
		// Early exit
		fmt.Println("Failure, early exit")
		os.Exit(1)
	}
	
	// Starting PickAxe service
	LoggerUtil.Logger.Info("Starting %s service", ConfigUtil.Get("service.name"))	
	var router *mux.Router = mux.NewRouter()
	router.HandleFunc("/healthcheck", healthcheck)
	var apiRouter *mux.Router = router.PathPrefix("/api/v1").Subrouter()
	SubscriberRouter.Register(apiRouter)
	log.Fatal(http.ListenAndServe(":8050", router))
}