package experimentation

import (
	"brownout-controller/constants"
	"brownout-controller/kubernetesCluster"
	"brownout-controller/powerModel"
	"brownout-controller/prometheus"
	"brownout-controller/util"
	"fmt"
	"log"
	"math"
	"time"
)

func HUCFExperiment(requiredSR float64) {

	sortedPods := kubernetesCluster.GetPodsSortedCPUUsageAllDescending(constants.NAMESPACE, constants.OPTIONAL)

	n := len(sortedPods)

	if n == 0 {
		return
	}

	m := n / 2 // mid point

	var i float64 = 0
	var podsToDeactivate []string
	var deactivatedPods map[string]int32

	for i < math.Log2(float64(n)) {

		fmt.Println("Iteration: ", i)

		currentSR := prometheus.GetSLASuccessRatio(constants.HOSTNAME, constants.SLA_INTERVAL, constants.SLA_VIOLATION_LATENCY)
		fmt.Println("Current SR: ", currentSR)

		if math.Abs(currentSR-requiredSR) < 0.05 {
			break
		} else if currentSR > requiredSR {
			if i != 0 {
				m = (m + n) / 2
				kubernetesCluster.ActivatePods(deactivatedPods, constants.NAMESPACE)
				time.Sleep(30 * time.Second)

				terminatingPods := kubernetesCluster.GetTerminatingPodNamesAll(constants.NAMESPACE)
				fmt.Println("Terminating Pods: ", terminatingPods)

				tempSlice := kubernetesCluster.GetPodsSortedCPUUsageAllDescending(constants.NAMESPACE, constants.OPTIONAL)
				fmt.Println("Temp Pods: ", tempSlice)

				sortedPods = util.SliceDifference(tempSlice, terminatingPods)
			}
			fmt.Println("m: ", m)
			podsToDeactivate = sortedPods[:m+1]
			deactivatedPods = kubernetesCluster.DeactivatePods(podsToDeactivate, constants.NAMESPACE)
		} else {
			if i != 0 {
				m = (1 + m) / 2
				kubernetesCluster.ActivatePods(deactivatedPods, constants.NAMESPACE)
				time.Sleep(30 * time.Second)

				terminatingPods := kubernetesCluster.GetTerminatingPodNamesAll(constants.NAMESPACE)
				fmt.Println("Terminating Pods: ", terminatingPods)

				tempSlice := kubernetesCluster.GetPodsSortedCPUUsageAllDescending(constants.NAMESPACE, constants.OPTIONAL)
				fmt.Println("Temp Pods: ", tempSlice)

				sortedPods = util.SliceDifference(tempSlice, terminatingPods)
			}
			fmt.Println("m: ", m)
			podsToDeactivate = sortedPods[:m+1]
			deactivatedPods = kubernetesCluster.DeactivatePods(podsToDeactivate, constants.NAMESPACE)
		}
		fmt.Println("Deactivated Pods: ", deactivatedPods)
		i++

		fmt.Println("Waiting for 5 minutes")
		time.Sleep(5 * time.Minute)
	}

	allClusterPods := kubernetesCluster.GetPodNamesAll(constants.NAMESPACE)
	predictedClusterPods := util.SliceDifference(allClusterPods, podsToDeactivate)

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
