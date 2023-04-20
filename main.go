package main

import (
	"brownout-controller/constants"
	"brownout-controller/kubernetesCluster"
	"brownout-controller/powerModel"
	"fmt"
)

func main() {

	//fmt.Println(prometheus.GetSLAViolationRatio(constants.HOSTNAME, constants.SLA_INTERVAL, constants.SLA_VIOLATION_LATENCY))
	//fmt.Println(prometheus.GetSLASuccessRatio(constants.HOSTNAME, constants.SLA_INTERVAL, constants.SLA_VIOLATION_LATENCY))
	pods := kubernetesCluster.GetPodNamesAll(constants.NAMESPACE)
	fmt.Println(powerModel.GetPowerModel().GetPowerConsumptionPods(pods))
}
