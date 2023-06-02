package experimentation

import (
	"brownout-controller/constants"
	"brownout-controller/kubernetesCluster"
	"brownout-controller/powerModel"
	"brownout-controller/prometheus"
	"brownout-controller/variables"
	"log"
)

// DoBrownoutExperimentNodePolicy Deactivate containers based on a container selection policy
// (LUCF, LRU, Random)
func DoBrownoutExperimentNodePolicy(policy string, policyK float64) {

	currentSuccessRate := prometheus.GetSLASuccessRatio(constants.HOSTNAME, variables.SLA_INTERVAL, variables.SLA_VIOLATION_LATENCY)
	log.Println("Initial SR: ", currentSuccessRate)

	// ACCEPTED_SUCCESS_RATE = approx. 0.65
	if currentSuccessRate > variables.ACCEPTED_SUCCESS_RATE {
		currentPowerConsumption := powerModel.GetPowerModel().GetPowerConsumptionPods(kubernetesCluster.GetPodNamesAll(constants.NAMESPACE))
		log.Println("Initial Power: ", currentPowerConsumption)

		// current_success_rate / ACCEPTED_SUCCESS_RATE = k * (current_power_consumption / upper_threshold_power )
		upperThresholdPower := policyK * (currentPowerConsumption * variables.ACCEPTED_SUCCESS_RATE / currentSuccessRate)
		log.Println("Policy K: ", policyK)
		log.Println("ASR: ", variables.ACCEPTED_SUCCESS_RATE)
		log.Println("Calculated upper threshold Power: ", upperThresholdPower)

		DoExperimentNodePolicies(policy, upperThresholdPower)
	}
}
