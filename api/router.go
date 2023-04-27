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

	// Route handles & endpoints
	// TODO Samples below. Delete later.
	s.HandleFunc("/books", getBooks).Methods("GET")

	// Websocket Sample
	s.HandleFunc("/ws", handleWebSocket)
}

func initBrownoutSubRouter(r *mux.Router) {
	s := r.PathPrefix("/brownout").Subrouter()

	// TODO Samples below. Delete later.
	s.HandleFunc("/books", getBooks).Methods("GET")
	s.HandleFunc("/books/{id}", getBook).Methods("GET")
	s.HandleFunc("/books", createBook).Methods("POST")
	s.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	s.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")
}
