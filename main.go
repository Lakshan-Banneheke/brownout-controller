package main

import "brownout-controller/api"

func main() {

	//pods := kubernetesCluster.GetPodNamesAll(constants.NAMESPACE)
	//fmt.Println("Initial Power: ", powerModel.GetPowerModel().GetPowerConsumptionPods(pods))
	//prometheus.GetSLASuccessRatio(constants.HOSTNAME, constants.SLA_INTERVAL, constants.SLA_VIOLATION_LATENCY)

	//nisp := policies.NISP{}
	////rcsp := policies.RCSP{}
	//
	////experimentation.DoBrownoutExperimentPodPolicy(rcsp, constants.K_RSCP)
	//experimentation.DoBrownoutExperimentNodePolicy(nisp, constants.K_NISP)

	//experimentation.DoExperimentPodPolicies(policy, 12)

	//fmt.Println(prometheus.GetSLAViolationRatio(constants.HOSTNAME, constants.SLA_INTERVAL, constants.SLA_VIOLATION_LATENCY))
	//fmt.Println(prometheus.GetSLASuccessRatio(constants.HOSTNAME, constants.SLA_INTERVAL, constants.SLA_VIOLATION_LATENCY))
	//pods := kubernetesCluster.GetPodNamesAll(constants.NAMESPACE)
	//for {
	//	fmt.Println("Current Power: ", powerModel.GetPowerModel().GetPowerConsumptionPods(pods))
	//	prometheus.GetSLASuccessRatio(constants.HOSTNAME, constants.SLA_INTERVAL, constants.SLA_VIOLATION_LATENCY)
	//	time.Sleep(30 * time.Second)
	//}

	api.InitAPI()
}
