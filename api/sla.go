package api

import (
	"brownout-controller/constants"
	"brownout-controller/prometheus"
	"log"
	"net/http"
	"time"
)

type SLAData struct {
	Timestamp     int64   `json:"timestamp"`
	TotReq        int     `json:"tot_req"`
	ErrReq        int     `json:"err_req"`
	SlowReq       int     `json:"slow_req"`
	TotSuccessReq int     `json:"tot_success_req"`
	SLASuccess    float64 `json:"sla_success"`
}

// User can use this API endpoint to get SLA related information
func handleListenSLA(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Client Connected to listen sla")
	for {
		slaData := SLAData{
			Timestamp:     time.Now().Unix(),
			TotReq:        prometheus.GetTotalRequestCount(constants.HOSTNAME, constants.SLA_INTERVAL),
			ErrReq:        prometheus.GetErrorRequestCount(constants.HOSTNAME, constants.SLA_INTERVAL),
			SlowReq:       prometheus.GetSlowRequestCount(constants.HOSTNAME, constants.SLA_INTERVAL, constants.SLA_VIOLATION_LATENCY),
			TotSuccessReq: prometheus.GetTotalSuccessRequestCount(constants.HOSTNAME, constants.SLA_INTERVAL),
			SLASuccess:    prometheus.GetSLASuccessRatio(constants.HOSTNAME, constants.SLA_INTERVAL, constants.SLA_VIOLATION_LATENCY),
		}

		// Send the data to the client
		err := conn.WriteJSON(slaData)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("SLA values written")

		// Wait for some time before sending the next data
		time.Sleep(30 * time.Second)
	}
}
