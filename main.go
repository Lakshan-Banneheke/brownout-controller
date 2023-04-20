package main

import (
	"brownout-controller/constants"
	"brownout-controller/kubernetesCluster"
	"brownout-controller/policies/experimentationv2"
	"brownout-controller/powerModel"
	"brownout-controller/prometheus"
	"fmt"
)

func main() {

	//thresholds := []float64{19, 18.75, 18.5, 18.25, 18, 17.5, 17, 16.5, 16}

	pods := kubernetesCluster.GetPodNamesAll(constants.NAMESPACE)
	fmt.Println("Initial Power: ", powerModel.GetPowerModel().GetPowerConsumptionPods(pods))
	prometheus.GetSLASuccessRatio(constants.HOSTNAME, constants.SLA_INTERVAL, constants.SLA_VIOLATION_LATENCY)

	experimentationv2.DoExperiment(18.75)

	//fmt.Println(prometheus.GetSLAViolationRatio(constants.HOSTNAME, constants.SLA_INTERVAL, constants.SLA_VIOLATION_LATENCY))
	//fmt.Println(prometheus.GetSLASuccessRatio(constants.HOSTNAME, constants.SLA_INTERVAL, constants.SLA_VIOLATION_LATENCY))
	//pods := kubernetesCluster.GetPodNamesAll(constants.NAMESPACE)
	//for {
	//	fmt.Println("Current Power: ", powerModel.GetPowerModel().GetPowerConsumptionPods(pods))
	//	prometheus.GetSLASuccessRatio(constants.HOSTNAME, constants.SLA_INTERVAL, constants.SLA_VIOLATION_LATENCY)
	//	time.Sleep(30 * time.Second)
	//}
}
