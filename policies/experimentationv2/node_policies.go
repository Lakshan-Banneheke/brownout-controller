package experimentationv2

import (
	"brownout-controller/constants"
	"brownout-controller/kubernetesCluster"
	"brownout-controller/policies"
	"brownout-controller/powerModel"
	"brownout-controller/prometheus"
	"brownout-controller/util"
	"fmt"
	"log"
	"time"
)

func DoExperimentNodePolicies(policy policies.IPolicyNodes, upperThresholdPower float64) {

	deactivatedPods, deactivatedNodes := policy.ExecuteForCluster(upperThresholdPower)
	log.Println("Deactivated Pods: ", deactivatedPods)
	log.Println("Deactivated Nodes: ", deactivatedNodes)

	log.Println("Waiting 3 minutes")
	time.Sleep(5 * time.Minute)

	allClusterPods := kubernetesCluster.GetPodNamesAll(constants.NAMESPACE)
	allNodes := kubernetesCluster.GetAllNodeNames()
	activeNodes := util.SliceDifference(allNodes, deactivatedNodes)

	log.Println("Pods after deactivation: ", allClusterPods)
	log.Println("Active nodes after deactivation: ", activeNodes)

	var predictedPowerList []float64
	var srList []float64

	fmt.Println("Getting power and SR")
	for i := 1; i <= 30; i++ {
		log.Println("==================================================================")

		srList = append(srList, prometheus.GetSLASuccessRatio(constants.HOSTNAME, constants.SLA_INTERVAL, constants.SLA_VIOLATION_LATENCY))
		// get power consumption of the pods
		predictedPowerList = append(predictedPowerList, powerModel.GetPowerModel().GetPowerConsumptionNodes(activeNodes))
		log.Println("Predicted Power List: ", predictedPowerList)
		log.Println("SR List: ", srList)

		avgPower := average(predictedPowerList)
		avgSr := average(srList)

		log.Println("Average SR: ", avgSr)
		log.Println("Average Power: ", avgPower)

		i++
		time.Sleep(1 * time.Second)
	}

	avgPower := average(predictedPowerList)
	avgSr := average(srList)

	log.Println("Average SR: ", avgSr)
	log.Println("Average Power: ", avgPower)

	log.Println("Upper threshold power: ", upperThresholdPower)
}
