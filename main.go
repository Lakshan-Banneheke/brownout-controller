package main

import (
	"brownout-controller/constants"
	"brownout-controller/kubernetesCluster"
	"brownout-controller/powerModel"
	"brownout-controller/prometheus"
	"fmt"
	"time"
)

func main() {

	//fmt.Println(prometheus.GetSLAViolationRatio(constants.HOSTNAME, constants.SLA_INTERVAL, constants.SLA_VIOLATION_LATENCY))
	//fmt.Println(prometheus.GetSLASuccessRatio(constants.HOSTNAME, constants.SLA_INTERVAL, constants.SLA_VIOLATION_LATENCY))
	pods := kubernetesCluster.GetPodNamesAll(constants.NAMESPACE)
	for {
		fmt.Println("Current Power: ", powerModel.GetPowerModel().GetPowerConsumptionPods(pods))
		prometheus.GetSLASuccessRatio(constants.HOSTNAME, constants.SLA_INTERVAL, constants.SLA_VIOLATION_LATENCY)
		time.Sleep(30 * time.Second)
	}
}
