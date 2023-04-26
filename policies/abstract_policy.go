package policies

import (
	"brownout-controller/constants"
	"brownout-controller/kubernetesCluster"
	"brownout-controller/powerModel"
	"brownout-controller/util"
	"log"
	"math"
)

type AbstractPolicy struct{}

func (absPolicy AbstractPolicy) executePolicy(allClusterPods []string, sortedPods []string, upperThresholdPower float64) map[string]int32 {

	n := len(sortedPods)
	if n == 0 {
		return make(map[string]int32)
	}

	min := 1
	max := n

	m := n / 2 // mid point

	var i float64 = 0
	var podsToDeactivate []string
	var predictedPower float64

	// performing a binary search to get the optimum cluster configuration
	for i < math.Log2(float64(n)) {
		log.Println("===============================================================")
		log.Println("Iteration: ", i)
		log.Println("m: ", m)
		podsToDeactivate = sortedPods[:m+1]

		// get the pods remaining in the cluster after deactivating above pods
		predictedClusterPods := util.SliceDifference(allClusterPods, podsToDeactivate)

		// get power consumption of the pods
		predictedPower = powerModel.GetPowerModel().GetPowerConsumptionPods(predictedClusterPods)
		log.Println("Predicted Power", predictedPower)
		log.Println("Upper Threshold", upperThresholdPower)

		if predictedPower > upperThresholdPower {
			min = m
			m = (m + max) / 2
		} else if (upperThresholdPower-predictedPower)/(upperThresholdPower) < 0.05 {
			break
		} else {
			max = m
			m = (min + m) / 2
		}

		i++
	}
	log.Println("Value for identified: ", m)
	log.Println("Deactivating pods")
	return kubernetesCluster.DeactivatePods(podsToDeactivate, constants.NAMESPACE)
}
