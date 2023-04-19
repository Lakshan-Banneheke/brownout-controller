package policies

import (
	"brownout-controller/constants"
	"brownout-controller/kubernetesCluster"
	"brownout-controller/powerModel"
)

// NISP implements the IPolicyNodes interface
type NISP struct{}

var node_deployments = make(map[string]int32)

// ExecuteForCluster
// Assumption: optional containers are deployed in nodes that are labelled as optional
// These nodes do not contain mandatory containers
func (nisp NISP) ExecuteForCluster(upperThresholdPower float64) map[string]int32 {
	sortedNodes := kubernetesCluster.GetNodesSortedCPUUsageAscending(constants.OPTIONAL)
	allNodes := sortedNodes
	return nisp.executePolicy(allNodes, sortedNodes, upperThresholdPower)
}

func (nisp NISP) executePolicy(allNodes []string, sortedNodes []string, upperThresholdPower float64) map[string]int32 {

	i := 0
	var predictedPower float64 = 0

	for predictedPower > upperThresholdPower {
		i++
		predictedClusterNodes := SliceDifference(allNodes, sortedNodes[0:i]) // get the nodes remaining in the cluster after deactivating i nodes

		// get power consumption of the nodes
		predictedPower = powerModel.GetPowerModel().GetPowerConsumptionNodes(predictedClusterNodes)
	}

	if (upperThresholdPower-predictedPower)/upperThresholdPower < 0.1 {
		nisp.deactivateNodes(sortedNodes[0:i]) // deactivate all pods of 0 to i hosts
		return node_deployments
	}

	var policy IPolicyPods = LUCF{} // Can set policy (LUCF, LRU, RCSP)

	if i == 1 {
		return policy.ExecuteForNode(sortedNodes[0], upperThresholdPower) // deactivate some containers of 0th node according to a pod selection policy
	} else {
		nisp.deactivateNodes(sortedNodes[0 : i-1])                                         // deactivate all containers of 0 to i-1 hosts
		one_node_deployments := policy.ExecuteForNode(sortedNodes[i], upperThresholdPower) // deactivate some containers of ith node according to a pod selection policy
		for key, value := range one_node_deployments {
			node_deployments[key] = value
		}

		return node_deployments
	}
}

func (nisp NISP) deactivateNodes(nodeList []string) {
	for _, node := range nodeList {
		one_node_deployments := kubernetesCluster.DeactivateNode(node, constants.NAMESPACE, constants.OPTIONAL)
		for key, value := range one_node_deployments {
			node_deployments[key] = value
		}

	}
}
