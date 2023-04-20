package experimentationv2

import (
	"brownout-controller/constants"
	"brownout-controller/kubernetesCluster"
	"brownout-controller/policies"
	"brownout-controller/powerModel"
	"brownout-controller/prometheus"
	"fmt"
	"log"
	"time"
)

func DoExperiment(upperThresholdPower float64) {

	lucf := policies.LUCF{}

	deactivatedPods := lucf.ExecuteForCluster(upperThresholdPower)
	log.Println("Deactivated Pods: ", deactivatedPods)

	log.Println("Waiting 5 minutes")
	time.Sleep(5 * time.Minute)

	allClusterPods := kubernetesCluster.GetPodNamesAll(constants.NAMESPACE)
	log.Println("Pods after deactivation: ", allClusterPods)

	var predictedPowerList []float64
	var srList []float64

	fmt.Println("Getting power and SR")
	for i := 1; i <= 60; i++ {
		// get power consumption of the pods
		predictedPowerList = append(predictedPowerList, powerModel.GetPowerModel().GetPowerConsumptionPods(allClusterPods))
		srList = append(srList, prometheus.GetSLASuccessRatio(constants.HOSTNAME, constants.SLA_INTERVAL, constants.SLA_VIOLATION_LATENCY))
		i++
		time.Sleep(1 * time.Second)
	}

	avgPower := average(predictedPowerList)
	avgSr := average(srList)

	log.Println("Upper threshold power: ", upperThresholdPower)
	log.Println("Average SR: ", avgSr)
	log.Println("Average Power: ", avgPower)
}
