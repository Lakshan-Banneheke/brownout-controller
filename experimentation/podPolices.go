package experimentation

import (
	"brownout-controller/constants"
	"brownout-controller/kubernetesCluster"
	"brownout-controller/policies"
	"brownout-controller/powerModel"
	"brownout-controller/prometheus"
	"brownout-controller/variables"
	"fmt"
	"log"
	"time"
)

func DoExperimentPodPolicies(policyName string, upperThresholdPower float64) {
	log.Printf("Running experiment for %s policy at Upper Threshold = %vW", policyName, upperThresholdPower)

	policy := policies.GetSelectedPolicy(policyName)

	prometheus.GetSLASuccessRatio(constants.HOSTNAME, variables.SLA_INTERVAL, variables.SLA_VIOLATION_LATENCY)

	deactivatedPods := policy.ExecuteForCluster(upperThresholdPower)
	log.Println("Deactivated Pods: ", deactivatedPods)

	log.Println("Waiting 3 minutes")
	time.Sleep(3 * time.Minute)

	allClusterPods := kubernetesCluster.GetPodNamesAll(constants.NAMESPACE)
	log.Println("Pods after deactivation: ", allClusterPods)

	var predictedPowerList []float64
	var srList []float64

	fmt.Println("Getting power and SR")
	for i := 1; i <= 30; i++ {
		log.Println("==================================================================")

		srList = append(srList, prometheus.GetSLASuccessRatio(constants.HOSTNAME, variables.SLA_INTERVAL, variables.SLA_VIOLATION_LATENCY))
		// get power consumption of the pods
		predictedPowerList = append(predictedPowerList, powerModel.GetPowerModel().GetPowerConsumptionPods(allClusterPods))
		log.Println("Predicted Power List: ", predictedPowerList)
		log.Println("SR List: ", srList)

		avgPower := average(predictedPowerList)
		avgSr := average(srList)

		log.Println("Average SR: ", avgSr)
		log.Println("Average Power: ", avgPower)

		i++
		time.Sleep(1 * time.Second)
	}

	avgPower := average(predictedPowerList)
	avgSr := average(srList)

	log.Println("Average SR: ", avgSr)
	log.Println("Average Power: ", avgPower)

	//log.Println("Upper threshold power: ", upperThresholdPower)
}
