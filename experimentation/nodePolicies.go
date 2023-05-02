package experimentation

import (
	"brownout-controller/constants"
	"brownout-controller/kubernetesCluster"
	"brownout-controller/policies"
	"brownout-controller/powerModel"
	"brownout-controller/prometheus"
	"log"
	"time"
)

func DoExperimentNodePolicies(policyName string, upperThresholdPower float64) {
	log.Printf("Running experiment for %s policy at Upper Threshold = %vW", policyName, upperThresholdPower)

	policy := policies.GetSelectedPolicy(policyName)

	prometheus.GetSLASuccessRatio(constants.HOSTNAME, constants.SLA_INTERVAL, constants.SLA_VIOLATION_LATENCY)

	deactivatedPods := policy.ExecuteForCluster(upperThresholdPower)
	log.Println("Deactivated Pods: ", deactivatedPods)

	log.Println("Waiting 4 minutes")
	time.Sleep(4 * time.Minute)

	activeNodes := kubernetesCluster.GetActiveNodeNames()

	log.Println("Active nodes after deactivation: ", activeNodes)

	var predictedPowerList []float64
	var srList []float64

	log.Println("Getting power and SR")
	for i := 1; i <= 30; i++ {
		log.Println("==================================================================")

		srList = append(srList, prometheus.GetSLASuccessRatio(constants.HOSTNAME, constants.SLA_INTERVAL, constants.SLA_VIOLATION_LATENCY))
		// get power consumption of the pods
		// NOTE: Don't use GetPowerConsumptionNodesWithMigration() here as migration has already happened
		predictedPowerList = append(predictedPowerList, powerModel.GetPowerModel().GetPowerConsumptionNodes(activeNodes))
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
	kubernetesCluster.UncordonAllNodes()
	//log.Println("Upper threshold power: ", upperThresholdPower)

}
