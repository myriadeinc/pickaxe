package Routes

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func subscriptionHandler (res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "OK")
}


func Register(router *mux.Router) {
	var subscriberRouter *mux.Router = router.Host("/subscribe").Subrouter()
	subscriberRouter.HandleFunc("/", subscriptionHandler).Methods("POST")
}