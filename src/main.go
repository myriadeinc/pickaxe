package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func healthcheck(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "OK")
}

func initializeService() bool {
	//Initialize our Logger
	LoggerUtil.Init()
	ConfigUtil.Init()
	MoneroApi.Init("http://0.0.0.0:8040")
	// Allow polling service to connect to monero api instead?
	PollingService.start()
	return true
}

func main() {

	var success bool = initializeService()
	if !success {
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
