package policies

import (
	"brownout-controller/constants"
	"brownout-controller/kubernetesCluster"
	"fmt"
)

// NISP implements the IPolicyNodes interface
type NISP struct{}

func (nisp NISP) ExecuteForCluster() {
	//# Assumption: optional containers are deployed in nodes that are labelled as optional
	//# These nodes do not contain mandatory containers
	sortedNodes := kubernetesCluster.GetNodesSortedCPUUsage(constants.OPTIONAL)
	allNodes := sortedNodes
	nisp.executePolicy(allNodes, sortedNodes)
}

func (nisp NISP) executePolicy(allNodes []string, sortedNodes []string) {

	i := 0
	var predictedPower float64 = 0

	for predictedPower > constants.UPPER_THRESHOLD_POWER {
		i++

		// get the nodes remaining in the cluster after deactivating i nodes
		predictedClusterNodes := SliceDifference(allNodes, sortedNodes[0:i])

		// TODO integrate with the powerModel package
		fmt.Println(predictedClusterNodes)
		//predictedPower = powerModel.getNodesPower(predictedClusterNodes)
	}

	if (constants.UPPER_THRESHOLD_POWER-predictedPower)/constants.UPPER_THRESHOLD_POWER < 0.1 {
		// deactivate all pods of 0 to i hosts
		nisp.deactivateNodes(sortedNodes[0:i])
		return
	}

	// Can set policy (LUCF, LRU, RCSP)
	var policy IPolicyPods = LUCF{}

	if i == 1 {
		// deactivate some containers of 0th node according to one of the pod selection policies
		policy.ExecuteForNode(sortedNodes[0])
	} else {
		// deactivate all containers of 0 to i-1 hosts
		nisp.deactivateNodes(sortedNodes[0 : i-1])

		// deactivate some containers of ith node according to one of the pod selection policies
		policy.ExecuteForNode(sortedNodes[i])
	}
}

func (nisp NISP) deactivateNodes(nodeList []string) {
	// TODO implement this function in kubernetesCluster package
	//for node := range nodeList {
	//	kubernetesCluster.DeactivateNode(node)
	//}
}
