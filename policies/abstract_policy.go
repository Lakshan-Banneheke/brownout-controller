package policies

import (
	"brownout-controller/constants"
	"brownout-controller/kubernetesCluster"
	"brownout-controller/powerModel"
	"math"
)

type AbstractPolicy struct{}

func (absPolicy AbstractPolicy) executePolicy(allClusterPods []string, sortedPods []string) {

	n := len(sortedPods)

	if n == 0 {
		return
	}

	m := n / 2 // mid point

	var i float64 = 0
	var podsToDeactivate []string
	var predictedPower float64

	// performing a binary search to get the optimum cluster configuration
	for i < math.Log2(float64(n)) {

		podsToDeactivate = sortedPods[:m+1]

		// get the pods remaining in the cluster after deactivating above pods
		predictedClusterPods := SliceDifference(allClusterPods, podsToDeactivate)

		// get power consumption of the pods
		predictedPower = powerModel.GetPowerModel("v4").GetPowerConsumptionPods(predictedClusterPods)

		if predictedPower > constants.UPPER_THRESHOLD_POWER {
			m = (m + n) / 2
		} else if (constants.UPPER_THRESHOLD_POWER-predictedPower)/(constants.UPPER_THRESHOLD_POWER) < 0.1 {
			break
		} else {
			m = (1 + m) / 2
		}

		i++
	}
	kubernetesCluster.DeactivatePods(podsToDeactivate, constants.NAMESPACE)
}