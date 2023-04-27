package api

import (
	"encoding/json"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	message := "Welcome to the BonE API. BonE is a comprehensive solution to achieving energy efficiency using the Brownout approach on the Edge"
	json.NewEncoder(w).Encode(message)
}
