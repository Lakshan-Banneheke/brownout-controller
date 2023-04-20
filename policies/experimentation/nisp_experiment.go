package experimentation

import (
	"brownout-controller/constants"
	"brownout-controller/kubernetesCluster"
	"brownout-controller/prometheus"
	"fmt"
	"math"
)

func NISPExperiment(requiredSR float64) {

	//Nodes sorted according to CPU util
	sortedNodes := kubernetesCluster.GetNodesSortedCPUUsageAscending(constants.OPTIONAL)

	//Number of nodes
	var i int32 = 0

	var nodesToDeactivate []string

	var podsToDeactivate []string
	var deactivatedPods map[string]int32

	for i < constants.OPTIONAL_NODES_NUM {
		i++
		nodesToDeactivate = sortedNodes[0:i]

		currentSR := prometheus.GetSLASuccessRatio(constants.HOSTNAME, constants.SLA_INTERVAL, constants.SLA_VIOLATION_LATENCY)
		fmt.Println("Current SR: ", currentSR)
		fmt.Println("deactivated nodes count: ", i)

		//if lower than min accepted rate - do not off nodes
		if math.Abs(currentSR-requiredSR) < 0.05 {
			break

		} else if (currentSR - requiredSR) > 0.75 {
			//if current sr is higher than the accepted sr - off more nodes
			i++
			nodesToDeactivate = sortedNodes[0:i]
			if i != 0 {
				//TODO
				kubernetesCluster.ActivatePods(deactivatedPods, constants.NAMESPACE)
				sortedNodes = kubernetesCluster.GetNodesSortedCPUUsageAscending(constants.OPTIONAL)
			}
			for _, node := range nodesToDeactivate {
				kubernetesCluster.DeactivateNode(node, constants.NAMESPACE, constants.OPTIONAL)
			}
		} else {
			// if current sr is between minimum accepted sr and accepted sr - off nodes + few pods
			nodesToDeactivate = sortedNodes[0:i]
			if i != 0 {
				//TODO
				kubernetesCluster.ActivatePods(deactivatedPods, constants.NAMESPACE)
				sortedNodes = kubernetesCluster.GetNodesSortedCPUUsageAscending(constants.OPTIONAL)
			}
			for _, node := range nodesToDeactivate {
				kubernetesCluster.DeactivateNode(node, constants.NAMESPACE, constants.OPTIONAL)
			}
			podsToDeactivate = kubernetesCluster.GetPodsInNode(sortedNodes[i], constants.NAMESPACE, constants.OPTIONAL)
			deactivatedPods = kubernetesCluster.DeactivatePods(podsToDeactivate, constants.OPTIONAL)

		}
	}
}
