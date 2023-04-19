package brownout

import (
	"brownout-controller/constants"
	"brownout-controller/kubernetesCluster"
	"brownout-controller/policies"
	"brownout-controller/powerModel"
	"brownout-controller/prometheus"
	"fmt"
	"time"
)

func ExecuteBrownout() {
	PowerModel := powerModel.GetPowerModel()

	var prevDeactivatedDeployments = make(map[string]int32)

	for {
		currentSuccessRate := prometheus.GetSLASuccessRatio(constants.HOSTNAME, constants.SLA_INTERVAL, constants.SLA_VIOLATION_LATENCY)

		// ACCEPTED_SUCCESS_RATE = approx. 0.75
		if currentSuccessRate > constants.ACCEPTED_SUCCESS_RATE {
			currentPowerConsumption := PowerModel.GetPowerConsumptionPods(kubernetesCluster.GetPodNames(constants.NAMESPACE, constants.OPTIONAL))
			// current_success_rate / ACCEPTED_SUCCESS_RATE = k * (current_power_consumption / upper_threshold_power )
			upperThresholdPower := constants.K_VALUE * (currentPowerConsumption * constants.ACCEPTED_SUCCESS_RATE / currentSuccessRate)

			// Deactivate containers based on a container selection policy
			// (Node Idling, LUCF, LRU, Random)
			lucf := policies.LUCF{}

			//DEACTIVATE_CONTAINERS(upperThresholdPower)
			prevDeactivatedDeployments = lucf.ExecuteForCluster(upperThresholdPower)

			// ACCEPTED_LOW_SUCCESS_RATE = approx. 0.50
		} else if currentSuccessRate < constants.ACCEPTED_MINIMUM_SUCCESS_RATE {
			// Activate all the deactivated containers
			if len(prevDeactivatedDeployments) == 0 {
				fmt.Println("There are no containers to activate")
			} else {
				kubernetesCluster.ActivatePods(prevDeactivatedDeployments, constants.NAMESPACE)
			}
		}

		time.Sleep(5 * time.Minute)
	}
}
