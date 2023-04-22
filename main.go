package main

import (
	"brownout-controller/policies"
	"brownout-controller/policies/experimentationv2"
)

func main() {

	//thresholds := []float64{19, 18.5, 18, 17.5, 17, 16.5, 16}

	//nisp := policies.NISP{}
	policy := policies.RCSP{}
	experimentationv2.DoExperimentPodPolicies(policy, 16)
	//experimentationv2.DoExperimentNodePolicies(nisp, 12)

	//fmt.Println(prometheus.GetSLAViolationRatio(constants.HOSTNAME, constants.SLA_INTERVAL, constants.SLA_VIOLATION_LATENCY))
	//fmt.Println(prometheus.GetSLASuccessRatio(constants.HOSTNAME, constants.SLA_INTERVAL, constants.SLA_VIOLATION_LATENCY))
	//pods := kubernetesCluster.GetPodNamesAll(constants.NAMESPACE)
	//for {
	//	fmt.Println("Current Power: ", powerModel.GetPowerModel().GetPowerConsumptionPods(pods))
	//	prometheus.GetSLASuccessRatio(constants.HOSTNAME, constants.SLA_INTERVAL, constants.SLA_VIOLATION_LATENCY)
	//	time.Sleep(30 * time.Second)
	//}
}
