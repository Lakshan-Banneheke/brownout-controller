package test

import (
	"brownout-controller/constants"
	"brownout-controller/kubernetesCluster"
	"brownout-controller/policies"
	"brownout-controller/powerModel"
	"brownout-controller/prometheus"
	"log"
	"math"
	"time"
)

func LUCFPolicyTest(requiredSR float64) {

	allClusterPods := kubernetesCluster.GetPodNames(constants.NAMESPACE, constants.OPTIONAL)
	sortedPods := kubernetesCluster.GetPodsSortedCPUUsageAllAscending(constants.NAMESPACE, constants.OPTIONAL)

	n := len(sortedPods)

	if n == 0 {
		return
	}

	m := n / 2 // mid point

	var i float64 = 0
	var podsToDeactivate []string
	var deactivatedPods map[string]int32

	for i < math.Log2(float64(n)) {

		podsToDeactivate = sortedPods[:m+1]

		currentSR := prometheus.GetSLAViolationRatio("podinfo.localdev.me", "1min", constants.SLA_VIOLATION_LATENCY)

		if math.Abs(currentSR-requiredSR) < 0.1 {
			break
		} else if currentSR > requiredSR {
			m = (m + n) / 2
			kubernetesCluster.ActivatePods(deactivatedPods, constants.NAMESPACE)
			sortedPods = kubernetesCluster.GetPodsSortedCPUUsageAllAscending(constants.NAMESPACE, constants.OPTIONAL)
			deactivatedPods = kubernetesCluster.DeactivatePods(podsToDeactivate, constants.NAMESPACE)
		} else {
			m = (1 + m) / 2
			kubernetesCluster.ActivatePods(deactivatedPods, constants.NAMESPACE)
			sortedPods = kubernetesCluster.GetPodsSortedCPUUsageAllAscending(constants.NAMESPACE, constants.OPTIONAL)
			deactivatedPods = kubernetesCluster.DeactivatePods(podsToDeactivate, constants.NAMESPACE)
		}

		i++
		time.Sleep(1 * time.Minute)
	}
	// get the pods remaining in the cluster after deactivating above pods
	predictedClusterPods := policies.SliceDifference(allClusterPods, podsToDeactivate)
	var predictedPowerList []float64
	var srList []float64

	for i := 1; i <= 300; i++ {
		// get power consumption of the pods
		predictedPowerList = append(predictedPowerList, powerModel.GetPowerModel().GetPowerConsumptionPods(predictedClusterPods))
		srList = append(srList, prometheus.GetSLAViolationRatio("podinfo.localdev.me", "1min", constants.SLA_VIOLATION_LATENCY))
		i++
		time.Sleep(1 * time.Second)
	}

	avgPower := average(predictedPowerList)
	avgSr := average(srList)

	log.Println(avgPower)
	log.Println(avgSr)
}

func average(listFloat []float64) float64 {
	sum := 0.0
	for _, x := range listFloat {
		sum += x
	}
	return sum / float64(len(listFloat))
}
