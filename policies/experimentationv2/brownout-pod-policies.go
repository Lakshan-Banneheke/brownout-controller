package experimentationv2

import (
	"brownout-controller/constants"
	"brownout-controller/kubernetesCluster"
	"brownout-controller/policies"
	"brownout-controller/powerModel"
	"brownout-controller/prometheus"
	"log"
)

// DoBrownoutExperimentPodPolicy Deactivate containers based on a container selection policy
// (LUCF, LRU, Random)
func DoBrownoutExperimentPodPolicy(policy policies.IPolicyPods, policyK float64) {

	currentSuccessRate := prometheus.GetSLASuccessRatio(constants.HOSTNAME, constants.SLA_INTERVAL, constants.SLA_VIOLATION_LATENCY)
	log.Println("Initial SR: ", currentSuccessRate)

	// ACCEPTED_SUCCESS_RATE = approx. 0.65
	if currentSuccessRate > constants.ACCEPTED_SUCCESS_RATE {
		currentPowerConsumption := powerModel.GetPowerModel().GetPowerConsumptionPods(kubernetesCluster.GetPodNamesAll(constants.NAMESPACE))
		log.Println("Initial Power: ", currentPowerConsumption)

		// current_success_rate / ACCEPTED_SUCCESS_RATE = k * (current_power_consumption / upper_threshold_power )
		upperThresholdPower := policyK * (currentPowerConsumption * constants.ACCEPTED_SUCCESS_RATE / currentSuccessRate)
		log.Println("Policy K: ", policyK)
		log.Println("ASR: ", constants.ACCEPTED_SUCCESS_RATE)
		log.Println("Calculated upper threshold Power: ", upperThresholdPower)

		DoExperimentPodPolicies(policy, upperThresholdPower)
	}
}
