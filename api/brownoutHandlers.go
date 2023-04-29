package api

import (
	"brownout-controller/brownout"
	"net/http"
)

func handleBrownoutActivation(w http.ResponseWriter, r *http.Request) {
	brownout.SetBrownoutActive(true)
	go brownout.ActivateBrownout() // The function will be executed in a new thread (goroutine)
	w.WriteHeader(http.StatusOK)
}

func handleBrownoutDeactivation(w http.ResponseWriter, r *http.Request) {
	brownout.SetBrownoutActive(false)
	go brownout.DeactivateBrownout() // The function will be executed in a new thread (goroutine)
	w.WriteHeader(http.StatusOK)
}
