package api

import (
	"brownout-controller/constants"
	"brownout-controller/kubernetesCluster"
	"log"
	"net/http"
	"time"
)

type DeploymentData struct {
	Timestamp      int64                          `json:"timestamp"`
	DeploymentList []kubernetesCluster.Deployment `json:"deployment_list"`
}

func handleListenDeploymentData(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Client Connected to listen deployment data")
	for {

		deploymentData := DeploymentData{
			Timestamp:      time.Now().Unix(),
			DeploymentList: kubernetesCluster.GetDeploymentsAll(constants.NAMESPACE),
		}

		// Send the data to the client
		err := conn.WriteJSON(deploymentData)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("Deployment data values written")

		// Wait for some time before sending the next data
		time.Sleep(30 * time.Second)
	}
}
