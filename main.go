package main

import "brownout-controller/api"

func main() {

	//pods := kubernetesCluster.GetPodNamesAll(constants.NAMESPACE)
	//fmt.Println("Initial Power: ", powerModel.GetPowerModel().GetPowerConsumptionPods(pods))
	//prometheus.GetSLASuccessRatio(constants.HOSTNAME, constants.SLA_INTERVAL, constants.SLA_VIOLATION_LATENCY)

	////experimentation.DoBrownoutExperimentPodPolicy("RCSP" constants.K_RCSP)
	//experimentation.DoBrownoutExperimentNodePolicy("NISP", constants.K_NISP)

	//experimentation.DoExperimentPodPolicies(policy, 12)

	//fmt.Println(prometheus.GetSLAViolationRatio(constants.HOSTNAME, constants.SLA_INTERVAL, constants.SLA_VIOLATION_LATENCY))
	//fmt.Println(prometheus.GetSLASuccessRatio(constants.HOSTNAME, constants.SLA_INTERVAL, constants.SLA_VIOLATION_LATENCY))
	//pods := kubernetesCluster.GetPodNamesAll(constants.NAMESPACE)
	//for {
	//	fmt.Println("Current Power: ", powerModel.GetPowerModel().GetPowerConsumptionPods(pods))
	//	prometheus.GetSLASuccessRatio(constants.HOSTNAME, constants.SLA_INTERVAL, constants.SLA_VIOLATION_LATENCY)
	//	time.Sleep(30 * time.Second)
	//}

	//kubernetesCluster.DeactivateNode("test-node-worker-3", constants.NAMESPACE, constants.OPTIONAL)
	//kubernetesCluster.DeactivateNode("test-node-worker-5", constants.NAMESPACE, constants.OPTIONAL)

	//kubernetesCluster.UncordonAllNodes()

	//log.Println(kubernetesCluster.GetActiveWorkerNodeCount())

	api.InitAPI()
}
