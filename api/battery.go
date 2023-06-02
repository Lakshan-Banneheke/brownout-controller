package api

import (
	"brownout-controller/brownout"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type responseBodyBattery struct {
	Battery int
}

// User can use this API endpoint to periodically send the battery percentage to the brownout controller
func handleSetBattery(w http.ResponseWriter, r *http.Request) {

	var body responseBodyBattery
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

type BatteryData struct {
	Timestamp int64 `json:"timestamp"`
	Battery   int   `json:"battery"`
}

func handleListenBattery(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Client Connected to listen battery")
	for {
		batteryData := BatteryData{
			Timestamp: time.Now().Unix(),
			Battery:   brownout.GetBatteryPercentage(),
		}

		// Send the data to the client
		err := conn.WriteJSON(batteryData)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("Battery percentage value written")
		log.Printf("Timestamp: %v, Battery: %vW", batteryData.Timestamp, batteryData.Battery)

		// Wait for some time before sending the next data
		time.Sleep(30 * time.Second)
	}
}
