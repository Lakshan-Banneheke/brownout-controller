package policies

import (
	"brownout-controller/constants"
	"brownout-controller/kubernetesCluster"
	"fmt"
)

// NISP implements the IPolicyNodes interface
type NISP struct{}

// ExecuteForCluster
// Assumption: optional containers are deployed in nodes that are labelled as optional
// These nodes do not contain mandatory containers
func (nisp NISP) ExecuteForCluster() {
	sortedNodes := kubernetesCluster.GetNodesSortedCPUUsage(constants.OPTIONAL)
	allNodes := sortedNodes
	nisp.executePolicy(allNodes, sortedNodes)
}

func (nisp NISP) executePolicy(allNodes []string, sortedNodes []string) {

	i := 0
	var predictedPower float64 = 0

	for predictedPower > constants.UPPER_THRESHOLD_POWER {
		i++
		predictedClusterNodes := SliceDifference(allNodes, sortedNodes[0:i]) // get the nodes remaining in the cluster after deactivating i nodes

		// TODO integrate with the powerModel package
		fmt.Println(predictedClusterNodes)
		//predictedPower = powerModel.getNodesPower(predictedClusterNodes)
	}

	if (constants.UPPER_THRESHOLD_POWER-predictedPower)/constants.UPPER_THRESHOLD_POWER < 0.1 {
		nisp.deactivateNodes(sortedNodes[0:i]) // deactivate all pods of 0 to i hosts
		return
	}

	var policy IPolicyPods = LUCF{} // Can set policy (LUCF, LRU, RCSP)

	if i == 1 {
		policy.ExecuteForNode(sortedNodes[0]) // deactivate some containers of 0th node according to a pod selection policy
	} else {
		nisp.deactivateNodes(sortedNodes[0 : i-1]) // deactivate all containers of 0 to i-1 hosts
		policy.ExecuteForNode(sortedNodes[i])      // deactivate some containers of ith node according to a pod selection policy
	}
}

func (nisp NISP) deactivateNodes(nodeList []string) {
	for _, node := range nodeList {
		kubernetesCluster.DeactivateNode(node, constants.NAMESPACE, constants.OPTIONAL)
	}
}
