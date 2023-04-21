package main

import (
	"brownout-controller/constants"
	"brownout-controller/kubernetesCluster"
	"brownout-controller/policies"
	"brownout-controller/policies/experimentationv2"
	"brownout-controller/powerModel"
	"brownout-controller/prometheus"
	"fmt"
)

func main() {

	//allNodes := kubernetesCluster.GetAllNodeNames()
	//log.Println(allNodes)
	//log.Println(len(allNodes))
	//log.Println(powerModel.GetPowerModel().GetPowerConsumptionNodes(allNodes))
	//pods := kubernetesCluster.GetPodNamesAll(constants.NAMESPACE)
	//log.Println("Initial Power: ", powerModel.GetPowerModel().GetPowerConsumptionPods(pods))

	//thresholds := []float64{19, 18.5, 18, 17.5, 17, 16.5, 16}

	pods := kubernetesCluster.GetPodNamesAll(constants.NAMESPACE)
	fmt.Println("Initial Power: ", powerModel.GetPowerModel().GetPowerConsumptionPods(pods))
	prometheus.GetSLASuccessRatio(constants.HOSTNAME, constants.SLA_INTERVAL, constants.SLA_VIOLATION_LATENCY)

	nisp := policies.NISP{}
	//hucf := policies.HUCF{}
	//experimentationv2.DoExperimentPodPolicies(hucf, 16)
	experimentationv2.DoExperimentNodePolicies(nisp, 14)

	//fmt.Println(prometheus.GetSLAViolationRatio(constants.HOSTNAME, constants.SLA_INTERVAL, constants.SLA_VIOLATION_LATENCY))
	//fmt.Println(prometheus.GetSLASuccessRatio(constants.HOSTNAME, constants.SLA_INTERVAL, constants.SLA_VIOLATION_LATENCY))
	//pods := kubernetesCluster.GetPodNamesAll(constants.NAMESPACE)
	//for {
	//	fmt.Println("Current Power: ", powerModel.GetPowerModel().GetPowerConsumptionPods(pods))
	//	prometheus.GetSLASuccessRatio(constants.HOSTNAME, constants.SLA_INTERVAL, constants.SLA_VIOLATION_LATENCY)
	//	time.Sleep(30 * time.Second)
	//}
}
