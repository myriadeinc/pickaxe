package SubscriberRouter

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/myriadeinc/pickaxe/src/util/logger"
	"github.com/myriadeinc/pickaxe/src/services/subscriber"
)

type SubscribeRequest struct {
	Hostname 	*string	`json:"hostname"`
}

func subscriptionHandler (res http.ResponseWriter, req *http.Request) {
	// Performs a Request decoding + validation for fields and types
	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()
	var subscribeRequest SubscribeRequest
	err := decoder.Decode(&subscribeRequest)
	if err != nil {
		LoggerUtil.Logger.Error(err.Error())
    http.Error(res, "Invalid request format", http.StatusBadRequest)
    return
	}
	if subscribeRequest.Hostname == nil {
		LoggerUtil.Logger.Error("Missing `hostname` field in request")
		http.Error(res, "Invalid request format", http.StatusBadRequest)
    return
	}
	LoggerUtil.Logger.Info("Received subscribe request from " + (*subscribeRequest.Hostname))
	// Performs a call on Subscriber service to register this new subscriber
	SubscriberService.Subscribe((*subscribeRequest.Hostname))
	fmt.Fprintf(res, "OK")
}


func Register(router *mux.Router) {
	var subscriberRouter *mux.Router = router.PathPrefix("/subscribe").Subrouter()
	subscriberRouter.HandleFunc("/", subscriptionHandler).Methods("POST")
	LoggerUtil.Logger.Info("Register Subscriber Router")
}