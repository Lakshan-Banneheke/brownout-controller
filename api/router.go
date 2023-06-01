package api

import (
	"brownout-controller/constants"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// InitAPI Initializes the router and starts listening on the specified port
func InitAPI() {
	addr := ":" + constants.PORT
	log.Println("Initializing the API Server")
	log.Fatal(http.ListenAndServe(addr, initRouter()))
}

func initRouter() *mux.Router {
	// Init router
	r := mux.NewRouter()

	r.HandleFunc("/", home).Methods("GET")

	initMetricsSubRouter(r)
	initBrownoutSubRouter(r)
	return r
}

func initMetricsSubRouter(r *mux.Router) {
	s := r.PathPrefix("/metrics").Subrouter()

	s.HandleFunc("/power", handleListenPower)
	s.HandleFunc("/battery", handleListenBattery)
}

func initBrownoutSubRouter(r *mux.Router) {
	s := r.PathPrefix("/brownout").Subrouter()

	// Route handles & endpoints
	s.HandleFunc("/activate", handleBrownoutActivation).Methods("POST")
	s.HandleFunc("/deactivate", handleBrownoutDeactivation).Methods("POST")
	s.HandleFunc("/battery/set", handleSetBattery).Methods("POST")
}
