package api

import (
	"brownout-controller/kubernetesCluster"
	"log"
	"net/http"
	"time"
)

type NodeData struct {
	Timestamp   int64    `json:"timestamp"`
	AllNodes    []string `json:"nodes_all"`
	ActiveNodes []string `json:"nodes_active"`
}

func handleListenNodeData(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Client Connected to listen node data")
	for {
		nodeData := NodeData{
			Timestamp:   time.Now().Unix(),
			AllNodes:    kubernetesCluster.GetAllNodeNames(),
			ActiveNodes: kubernetesCluster.GetActiveNodeNames(),
		}

		// Send the data to the client
		err := conn.WriteJSON(nodeData)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("Node data values written")

		// Wait for some time before sending the next data
		time.Sleep(30 * time.Second)
	}
}
