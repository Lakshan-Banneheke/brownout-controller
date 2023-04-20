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

func NISPExperiment(requiredSR float64) {

	//Nodes sorted according to CPU util
	sortedNodes := kubernetesCluster.GetNodesSortedCPUUsageAscending(constants.OPTIONAL)

	//Number of nodes
	var i int32 = 0

	//var nodesToDeactivate []string
	var deactivatedNodes []string

	//var podsToDeactivate []string
	var tempDeactivatedPods map[string]int32
	var deactivatedPods map[string]int32

	for i < constants.OPTIONAL_NODES_NUM {

		currentSR := prometheus.GetSLASuccessRatio(constants.HOSTNAME, constants.SLA_INTERVAL, constants.SLA_VIOLATION_LATENCY)
		fmt.Println("Current SR: ", currentSR)
		fmt.Println("deactivated nodes count: ", i)

		//if lower than min accepted rate - do not off nodes
		if math.Abs(currentSR-requiredSR) < 0.05 {
			break

		} else if currentSR > requiredSR {
			tempDeactivatedPods = kubernetesCluster.DeactivateNode(sortedNodes[i], constants.NAMESPACE, constants.OPTIONAL)
			updateDeactivatedPods(tempDeactivatedPods, deactivatedPods)
			i++
		} else {
			break
		}

		time.Sleep(5 * time.Minute)

	}

	totalClusterNodes := kubernetesCluster.GetAllNodeNames()
	if i > 0 {
		deactivatedNodes = sortedNodes[0 : i-1]
	}

	predictedClusterWorkerNodes := util.SliceDifference(totalClusterNodes, deactivatedNodes)

	var predictedPowerList []float64
	var srList []float64

	fmt.Println("Exited Loop")
	for i := 1; i <= 300; i++ {
		// get power consumption of the nodes
		predictedPowerList = append(predictedPowerList, powerModel.GetPowerModel().GetPowerConsumptionNodes(predictedClusterWorkerNodes))
		srList = append(srList, prometheus.GetSLASuccessRatio(constants.HOSTNAME, constants.SLA_INTERVAL, constants.SLA_VIOLATION_LATENCY))
		i++
		time.Sleep(1 * time.Second)
	}

	avgPower := average(predictedPowerList)
	avgSr := average(srList)

	log.Println("Required SR: ", requiredSR)
	log.Println("Number of active worker nodes: ", len(predictedClusterWorkerNodes))
	log.Println("Average Power: ", avgPower)
	log.Println("Average SR: ", avgSr)

}

func updateDeactivatedPods(tempDeactivatedPods map[string]int32, deactivatedPods map[string]int32) {
	for key, value1 := range tempDeactivatedPods {
		if value2, exists := deactivatedPods[key]; exists {
			deactivatedPods[key] = value2 + value1
		} else {
			deactivatedPods[key] = value1
		}
	}
}

//
//		} else if currentSR > requiredSR {
//			// next node in sorted Nodes should be deactivated
//			time.Sleep(5 * time.Minute)
//
//			//if current sr is higher than the accepted sr - off more nodes
//			//i++
//			//if i != 0 {
//			//	kubernetesCluster.ActivatePods(deactivatedPods, constants.NAMESPACE)
//			//	time.Sleep(30 * time.Second)
//			//
//			//	sortedNodes = kubernetesCluster.GetNodesSortedCPUUsageAscending(constants.OPTIONAL)
//			//}
//			//nodesToDeactivate = sortedNodes[0:i]
//			//
//			//deactivatedPods = map[string]int32{}
//			//for _, node := range nodesToDeactivate {
//			//	tempDeactivatedPods := kubernetesCluster.DeactivateNode(node, constants.NAMESPACE, constants.OPTIONAL)
//			//	updateDeactivatedPods(tempDeactivatedPods, deactivatedPods)
//			//}
//		} else {
//			break
//		}
//		i++
//		//else {
//		//	// if current sr is between minimum accepted sr and accepted sr - off nodes + few pods
//		//	if i != 0 {
//		//		kubernetesCluster.ActivatePods(deactivatedPods, constants.NAMESPACE)
//		//		time.Sleep(30 * time.Second)
//		//
//		//		sortedNodes = kubernetesCluster.GetNodesSortedCPUUsageAscending(constants.OPTIONAL)
//		//	}
//		//	nodesToDeactivate = sortedNodes[0:i]
//		//
//		//	deactivatedPods = map[string]int32{}
//		//	for _, node := range nodesToDeactivate {
//		//		tempDeactivatedPods := kubernetesCluster.DeactivateNode(node, constants.NAMESPACE, constants.OPTIONAL)
//		//		updateDeactivatedPods(tempDeactivatedPods, deactivatedPods)
//		//	}
//		//	podsToDeactivate = kubernetesCluster.GetPodsInNodeCategory(sortedNodes[i], constants.NAMESPACE, constants.OPTIONAL)
//		//	tempDeactivatedPods := kubernetesCluster.DeactivatePods(podsToDeactivate, constants.OPTIONAL)
//		//	updateDeactivatedPods(tempDeactivatedPods, deactivatedPods)
//		//
//		//}
//	}
//}
//
//func updateDeactivatedPods(tempDeactivatedPods map[string]int32, deactivatedPods map[string]int32) {
//	for key, value1 := range tempDeactivatedPods {
//		if value2, exists := deactivatedPods[key]; exists {
//			deactivatedPods[key] = value2 + value1
//		} else {
//			deactivatedPods[key] = value1
//		}
//	}
//}
