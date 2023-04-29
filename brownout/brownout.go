package brownout

import (
	"brownout-controller/constants"
	"brownout-controller/kubernetesCluster"
	"brownout-controller/policies"
	"brownout-controller/powerModel"
	"brownout-controller/prometheus"
	"brownout-controller/util"
	"log"
	"time"
)

var deactivatedDeployments = map[string]int32{}

// ActivateBrownout is triggered if the signal is sent to activate the brownout algorithm through the API
func ActivateBrownout() {
	log.Println("Brownout has been activated")
	for {
		log.Printf("Checking battery percentage. Battery is at %v%%", batteryPercentage)
		if batteryPercentage < constants.BATTERY_LOWER_THRESHOLD {
			log.Printf("Battery percentage less than %v%%. Executing Brownout in the cluster", constants.BATTERY_LOWER_THRESHOLD)
			runBrownout()
		} else if batteryPercentage > constants.BATTERY_UPPER_THRESHOLD {
			log.Printf("Battery percentage greater than %v%%. Stopping Brownout in the cluster", constants.BATTERY_UPPER_THRESHOLD)
			stopBrownout()
		}
		time.Sleep(5 * time.Minute)
	}
}

// DeactivateBrownout is triggered if the signal is sent to activate the brownout algorithm through the API
func DeactivateBrownout() {
	log.Println("Brownout has been deactivated")
	stopBrownout()
}

func runBrownout() {
	currentSuccessRate := prometheus.GetSLASuccessRatio(constants.HOSTNAME, constants.SLA_INTERVAL, constants.SLA_VIOLATION_LATENCY)
	log.Println("Initial SR: ", currentSuccessRate)
	log.Println("ASR: ", constants.ACCEPTED_SUCCESS_RATE)

	// ACCEPTED_SUCCESS_RATE = approx. 0.65
	if currentSuccessRate > constants.ACCEPTED_SUCCESS_RATE {
		currentPowerConsumption := powerModel.GetPowerModel().GetPowerConsumptionPods(kubernetesCluster.GetPodNamesAll(constants.NAMESPACE))
		log.Println("Initial Power: ", currentPowerConsumption)

		// current_success_rate / ACCEPTED_SUCCESS_RATE = k * (current_power_consumption / upper_threshold_power )
		upperThresholdPower := constants.K_VALUE * (currentPowerConsumption * constants.ACCEPTED_SUCCESS_RATE / currentSuccessRate)
		log.Println("Calculated upper threshold Power: ", upperThresholdPower)

		// Deactivate containers based on a container selection policy
		// (Node Idling, LUCF, LRU, Random)
		lucf := policies.LUCF{}

		//DEACTIVATE_CONTAINERS(upperThresholdPower)
		currentDeactivatedDeployments := lucf.ExecuteForCluster(upperThresholdPower)
		deactivatedDeployments = util.AddToDeployments(currentDeactivatedDeployments, deactivatedDeployments)

		// ACCEPTED_LOW_SUCCESS_RATE = approx. 0.50
	} else if currentSuccessRate < constants.ACCEPTED_MINIMUM_SUCCESS_RATE {
		stopBrownout()
	}
}

func stopBrownout() {
	// Activate all the deactivated containers
	if len(deactivatedDeployments) == 0 {
		log.Println("There are no containers to activate")
	} else {
		kubernetesCluster.ActivatePods(deactivatedDeployments, constants.NAMESPACE)
		deactivatedDeployments = map[string]int32{}
	}
}
