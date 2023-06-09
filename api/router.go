package api

import (
	"brownout-controller/constants"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// InitAPI Initializes the router and starts listening on the specified port
func InitAPI() {
	addr := ":" + constants.PORT
	log.Println("Initializing the API Server")

	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),            // Allow requests from any origin
		handlers.AllowedMethods([]string{"GET", "POST"}),  // Allow GET and POST methods
		handlers.AllowedHeaders([]string{"Content-Type"}), // Allow "Content-Type" header
	)

	log.Fatal(http.ListenAndServe(addr, corsHandler(initRouter())))
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
	s.HandleFunc("/sla", handleListenSLA)
	s.HandleFunc("/nodes/list", handleListenNodeData)
	s.HandleFunc("/pods", handleListenPodData)
	s.HandleFunc("/deployments", handleListenDeploymentData)
}

func initBrownoutSubRouter(r *mux.Router) {
	s := r.PathPrefix("/brownout").Subrouter()

	// Route handles & endpoints
	s.HandleFunc("/activate", handleBrownoutActivation).Methods("POST")
	s.HandleFunc("/deactivate", handleBrownoutDeactivation).Methods("POST")
	s.HandleFunc("/battery/set", handleSetBattery).Methods("POST")
	s.HandleFunc("/variables/{name}", handleGetVariable).Methods("GET")
	s.HandleFunc("/variables/{name}", handleSetVariable).Methods("POST")
}
