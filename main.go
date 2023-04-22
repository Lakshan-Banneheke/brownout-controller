package main

import (
	"brownout-controller/constants"
	"brownout-controller/policies"
	"brownout-controller/policies/experimentationv2"
)

func main() {

	//pods := kubernetesCluster.GetPodNamesAll(constants.NAMESPACE)
	//fmt.Println("Initial Power: ", powerModel.GetPowerModel().GetPowerConsumptionPods(pods))
	//prometheus.GetSLASuccessRatio(constants.HOSTNAME, constants.SLA_INTERVAL, constants.SLA_VIOLATION_LATENCY)

	//nisp := policies.NISP{}
	lucf := policies.LUCF{}

	experimentationv2.DoBrownoutExperimentPodPolicy(lucf, constants.K_LUCF)

	//experimentationv2.DoExperimentPodPolicies(policy, 12)

	//fmt.Println(prometheus.GetSLAViolationRatio(constants.HOSTNAME, constants.SLA_INTERVAL, constants.SLA_VIOLATION_LATENCY))
	//fmt.Println(prometheus.GetSLASuccessRatio(constants.HOSTNAME, constants.SLA_INTERVAL, constants.SLA_VIOLATION_LATENCY))
	//pods := kubernetesCluster.GetPodNamesAll(constants.NAMESPACE)
	//for {
	//	fmt.Println("Current Power: ", powerModel.GetPowerModel().GetPowerConsumptionPods(pods))
	//	prometheus.GetSLASuccessRatio(constants.HOSTNAME, constants.SLA_INTERVAL, constants.SLA_VIOLATION_LATENCY)
	//	time.Sleep(30 * time.Second)
	//}
}
