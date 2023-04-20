package policies

import (
	"brownout-controller/constants"
	"brownout-controller/kubernetesCluster"
	"brownout-controller/powerModel"
	"brownout-controller/util"
	"math"
)

type AbstractPolicy struct{}

func (absPolicy AbstractPolicy) executePolicy(allClusterPods []string, sortedPods []string, upperThresholdPower float64) map[string]int32 {

	n := len(sortedPods)

	if n == 0 {
		return make(map[string]int32)
	}

	m := n / 2 // mid point

	var i float64 = 0
	var podsToDeactivate []string
	var predictedPower float64

	// performing a binary search to get the optimum cluster configuration
	for i < math.Log2(float64(n)) {

		podsToDeactivate = sortedPods[:m+1]

		// get the pods remaining in the cluster after deactivating above pods
		predictedClusterPods := util.SliceDifference(allClusterPods, podsToDeactivate)

		// get power consumption of the pods
		predictedPower = powerModel.GetPowerModel().GetPowerConsumptionPods(predictedClusterPods)

		if predictedPower > upperThresholdPower {
			m = (m + n) / 2
		} else if (upperThresholdPower-predictedPower)/(upperThresholdPower) < 0.1 {
			break
		} else {
			m = (1 + m) / 2
		}

		i++
	}
	return kubernetesCluster.DeactivatePods(podsToDeactivate, constants.NAMESPACE)
}
