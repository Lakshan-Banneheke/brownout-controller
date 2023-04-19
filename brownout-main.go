package main

import (
	"brownout-controller/constants"
	"brownout-controller/kubernetesCluster"
	"brownout-controller/policies"
	"brownout-controller/powerModel"
	"brownout-controller/prometheus"
	"fmt"
	"time"
)

func brownout() {
	PowerModel := powerModel.GetPowerModel(constants.POWER_MODEL_VERSION)

	var prev_deactivated_deployments = make(map[string]int32)

	for {
		currentSuccessRate := prometheus.GetSLASuccessRatio("podinfo.localdev.me", "1d", constants.SLA_VIOLATION_LATENCY)

		// ACCEPTED_SUCCESS_RATE = approx. 0.75
		if currentSuccessRate > constants.ACCEPTED_SUCCESS_RATE {
			currentPowerConsumption := PowerModel.GetPowerConsumptionPods(kubernetesCluster.GetPodNames(constants.NAMESPACE, ""))
			// current_success_rate / ACCEPTED_SUCCESS_RATE = k * (current_power_consumtpion / upper_threshold_power )
			upperThresholdPower := constants.K_VALUE * (currentPowerConsumption * constants.ACCEPTED_SUCCESS_RATE / currentSuccessRate)

			// Deactivate containers based on a container selection policy
			// (Node Idling, LUCF, LRU, Random)
			lucf := policies.LUCF{}

			//DEACTIVATE_CONTAINERS(upperThresholdPower)
			prev_deactivated_deployments = lucf.ExecuteForCluster(upperThresholdPower)

			// ACCEPTED_LOW_SUCCESS_RATE = approx. 0.50
		} else if currentSuccessRate < constants.ACCEPTED_MINIMUM_SUCCESS_RATE {
			// Activate all the deactivated containers
			if len(prev_deactivated_deployments) == 0 {
				fmt.Println("There are no containers to activate")
			} else {
				kubernetesCluster.ActivatePods(prev_deactivated_deployments, constants.NAMESPACE)
			}
		}

		time.Sleep(5 * time.Minute)
	}
}
