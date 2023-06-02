package api

import (
	"brownout-controller/brownout"
	"brownout-controller/variables"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
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

// Get a variable value
func handleGetVariable(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params

	name := params["name"]
	var value any

	switch name {
	case "policy":
		value = variables.POLICY
	case "batteryUpper":
		value = variables.BATTERY_UPPER_THRESHOLD
	case "batteryLower":
		value = variables.BATTERY_LOWER_THRESHOLD
	case "slaViolationLatency":
		value = variables.SLA_VIOLATION_LATENCY
	case "slaInterval":
		value = variables.SLA_INTERVAL
	case "asr":
		value = variables.ACCEPTED_SUCCESS_RATE
	case "amsr":
		value = variables.ACCEPTED_MIN_SUCCESS_RATE
	default:
		message := "Invalid variable name."
		http.Error(w, message, http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(value)
}

type responseBodyVariable struct {
	Value string
}

// Set a variable value
func handleSetVariable(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) // Gets params
	name := params["name"]

	var body responseBodyVariable
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		message := "Error in request body. Value must be of type string. Example request body: {\"Value\":\"45\"}"
		http.Error(w, message, http.StatusBadRequest)
		return
	}

	strValue := body.Value

	switch name {
	case "policy":
		variables.POLICY = strValue
	case "batteryUpper":
		value, err := strconv.Atoi(strValue)
		if err != nil {
			message := "Error in request body. Cannot convert value to integer. Example request body: {\"Value\":\"45\"}"
			http.Error(w, message, http.StatusBadRequest)
			return
		}
		variables.BATTERY_UPPER_THRESHOLD = value
	case "batteryLower":
		value, err := strconv.Atoi(strValue)
		if err != nil {
			message := "Error in request body. Cannot convert value to integer. Example request body: {\"Value\":\"45\"}"
			http.Error(w, message, http.StatusBadRequest)
			return
		}
		variables.BATTERY_LOWER_THRESHOLD = value
	case "slaViolationLatency":
		variables.SLA_VIOLATION_LATENCY = strValue
	case "slaInterval":
		variables.SLA_INTERVAL = strValue
	case "asr":
		value, err := strconv.ParseFloat(strValue, 64)
		if err != nil {
			message := "Error in request body. Cannot convert value to float. Example request body: {\"Value\":\"0.65\"}"
			http.Error(w, message, http.StatusBadRequest)
			return
		}
		variables.ACCEPTED_SUCCESS_RATE = value
	case "amsr":
		value, err := strconv.ParseFloat(strValue, 64)
		if err != nil {
			message := "Error in request body. Cannot convert value to float. Example request body: {\"Value\":\"0.65\"}"
			http.Error(w, message, http.StatusBadRequest)
			return
		}
		variables.ACCEPTED_MIN_SUCCESS_RATE = value
	default:
		message := "Error in request URL. Name of variable incorrect."
		http.Error(w, message, http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
