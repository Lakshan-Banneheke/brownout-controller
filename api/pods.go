package api

import (
	"brownout-controller/constants"
	"brownout-controller/kubernetesCluster"
	"log"
	"net/http"
	"time"
)

type PodData struct {
	Timestamp int64    `json:"timestamp"`
	PodList   []string `json:"pod_list"`
	Count     int      `json:"count"`
	CpuTotal  float64  `json:"cpu_total"`
}

func handleListenPodData(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	pods := kubernetesCluster.GetPodNamesAll(constants.NAMESPACE)
	podCPU := kubernetesCluster.GetPodsCPUUsageSum(pods, constants.NAMESPACE)

	log.Println("Client Connected to listen pod data")
	for {
		podData := PodData{
			Timestamp: time.Now().Unix(),
			PodList:   pods,
			Count:     len(pods),
			CpuTotal:  podCPU,
		}

		// Send the data to the client
		err := conn.WriteJSON(podData)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("Pod data values written")

		// Wait for some time before sending the next data
		time.Sleep(30 * time.Second)
	}
}
