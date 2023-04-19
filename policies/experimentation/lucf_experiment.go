package experimentation

import (
	"brownout-controller/constants"
	"brownout-controller/kubernetesCluster"
	"brownout-controller/policies"
	"brownout-controller/powerModel"
	"brownout-controller/prometheus"
	"fmt"
	"log"
	"math"
	"time"
)

func LUCFExperiment(requiredSR float64) {

	allClusterPods := kubernetesCluster.GetPodNamesAll(constants.NAMESPACE)
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

		currentSR := prometheus.GetSLASuccessRatio(constants.HOSTNAME, constants.SLA_INTERVAL, constants.SLA_VIOLATION_LATENCY)
		fmt.Println("Current SR: ", currentSR)
		fmt.Println("m: ", m)

		if math.Abs(currentSR-requiredSR) < 0.05 {
			break
		} else if currentSR > requiredSR {
			m = (m + n) / 2
			if i != 0 {
				kubernetesCluster.ActivatePods(deactivatedPods, constants.NAMESPACE)
				sortedPods = kubernetesCluster.GetPodsSortedCPUUsageAllAscending(constants.NAMESPACE, constants.OPTIONAL)
			}
			deactivatedPods = kubernetesCluster.DeactivatePods(podsToDeactivate, constants.NAMESPACE)
		} else {
			m = (1 + m) / 2
			if i != 0 {
				kubernetesCluster.ActivatePods(deactivatedPods, constants.NAMESPACE)
				sortedPods = kubernetesCluster.GetPodsSortedCPUUsageAllAscending(constants.NAMESPACE, constants.OPTIONAL)
			}
			deactivatedPods = kubernetesCluster.DeactivatePods(podsToDeactivate, constants.NAMESPACE)
		}

		i++
		time.Sleep(30 * time.Second)
	}
	// get the pods remaining in the cluster after deactivating above pods
	predictedClusterPods := policies.SliceDifference(allClusterPods, podsToDeactivate)
	var predictedPowerList []float64
	var srList []float64

	fmt.Println("Exited Loop")
	for i := 1; i <= 300; i++ {
		// get power consumption of the pods
		predictedPowerList = append(predictedPowerList, powerModel.GetPowerModel().GetPowerConsumptionPods(predictedClusterPods))
		srList = append(srList, prometheus.GetSLASuccessRatio(constants.HOSTNAME, constants.SLA_INTERVAL, constants.SLA_VIOLATION_LATENCY))
		i++
		time.Sleep(1 * time.Second)
	}

	avgPower := average(predictedPowerList)
	avgSr := average(srList)

	log.Println("Required SR: ", requiredSR)
	log.Println("Number of pods deactivated: ", len(podsToDeactivate))
	log.Println("Average Power: ", avgPower)
	log.Println("Average SR: ", avgSr)
}

func average(listFloat []float64) float64 {
	sum := 0.0
	for _, x := range listFloat {
		sum += x
	}
	return sum / float64(len(listFloat))
}