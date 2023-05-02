package main

import (
	"brownout-controller/experimentation"
	"os"
	"strconv"
)

func main() {

	policy := os.Getenv("policy") // get policy name from env variables eg: NISP
	uts := os.Getenv("UT")        // get upper threshold from env variables
	ut, _ := strconv.ParseFloat(uts, 64)
	experimentation.DoExperimentNodePolicies(policy, ut)

	//pods := kubernetesCluster.GetPodNamesAll(constants.NAMESPACE)
	//fmt.Println("Initial Power: ", powerModel.GetPowerModel().GetPowerConsumptionPods(pods))
	//prometheus.GetSLASuccessRatio(constants.HOSTNAME, constants.SLA_INTERVAL, constants.SLA_VIOLATION_LATENCY)

	////experimentation.DoBrownoutExperimentPodPolicy("RCSP" constants.K_RCSP)
	//experimentation.DoBrownoutExperimentNodePolicy("NISP", constants.K_NISP)

	//fmt.Println(prometheus.GetSLAViolationRatio(constants.HOSTNAME, constants.SLA_INTERVAL, constants.SLA_VIOLATION_LATENCY))
	//fmt.Println(prometheus.GetSLASuccessRatio(constants.HOSTNAME, constants.SLA_INTERVAL, constants.SLA_VIOLATION_LATENCY))
	//pods := kubernetesCluster.GetPodNamesAll(constants.NAMESPACE)
	//for {
	//	fmt.Println("Current Power: ", powerModel.GetPowerModel().GetPowerConsumptionPods(pods))
	//	prometheus.GetSLASuccessRatio(constants.HOSTNAME, constants.SLA_INTERVAL, constants.SLA_VIOLATION_LATENCY)
	//	time.Sleep(30 * time.Second)
	//}

}
