package api

import (
	"brownout-controller/brownout"
	"encoding/json"
	"net/http"
)

type responseBody struct {
	Battery int
}

// User can use this API endpoint to periodically send the battery percentage to the brownout controller
func handleSetBattery(w http.ResponseWriter, r *http.Request) {

	var body responseBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		message := "Battery Percentage should be an integer. Example request body: {\"Battery\":45}"
		http.Error(w, message, http.StatusBadRequest)
		return
	}

	batteryPercentage := body.Battery

	if batteryPercentage < 100 && batteryPercentage > 0 {
		brownout.SetBatteryPercentage(batteryPercentage)
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, "Battery Percentage should be an integer between 0 and 100", http.StatusBadRequest)
	}
}
