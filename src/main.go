package main

import (
	"os"
	"net/http"
	"fmt"
	"log"
	"github.com/gorilla/mux"
	"github.com/myriadeinc/pickaxe/src/api"
	"github.com/myriadeinc/pickaxe/src/util"
	"github.com/myriadeinc/pickaxe/src/routes"
	"github.com/myriadeinc/pickaxe/src/services"
)

func healthcheck(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "OK")
	var coolThing := Services.NewTemplateFetcher(9000)
}

func initializeService() (bool) {
	//Initialize our Logger
	Logger.Init()
	Apis.Init("http://0.0.0.0:8040")
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
	Logger.Logger.Info("Starting PickAxe service")
	var router *mux.Router = mux.NewRouter()
	router.HandleFunc("/healthcheck", healthcheck)
	var apiRouter *mux.Router = router.Host("/api/v1").Subrouter()
	SubscriberRouter.Register(apiRouter)
	log.Fatal(http.ListenAndServe(":8050", router))


	
}