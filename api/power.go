package api

import (
	"brownout-controller/powerModel"
	"log"
	"net/http"
	"time"
)

type PowerData struct {
	Timestamp int64   `json:"timestamp"`
	Power     float64 `json:"power"`
}

func handleGetPower(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Client Connected to GetPower")
	for {
		powerData := PowerData{
			Timestamp: time.Now().Unix(),
			Power:     powerModel.GetPowerModel().GetPowerConsumptionCluster(),
		}

		// Send the data to the client
		err := conn.WriteJSON(powerData)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("Power consumption value written")
		log.Printf("Timestamp: %v, Power: %vW", powerData.Timestamp, powerData.Power)

		// Wait for some time before sending the next data
		time.Sleep(1 * time.Second)
	}
}
