package main

import (
	"brownout-controller/constants"
	"brownout-controller/kubernetesCluster"
	"log"
)

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

	kubernetesCluster.DeactivateNode("test-node-worker-3", constants.NAMESPACE, constants.OPTIONAL)
	kubernetesCluster.DeactivateNode("test-node-worker-5", constants.NAMESPACE, constants.OPTIONAL)
	//
	//kubernetesCluster.UncordonAllNodes()

	log.Println(kubernetesCluster.GetActiveWorkerNodeCount())

}
