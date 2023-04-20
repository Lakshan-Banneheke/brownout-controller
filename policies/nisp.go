package policies

import (
	"brownout-controller/constants"
	"brownout-controller/kubernetesCluster"
	"brownout-controller/powerModel"
	"brownout-controller/util"
	"log"
)

// NISP implements the IPolicyNodes interface
type NISP struct{}

var nodeDeployments = make(map[string]int32)

// ExecuteForCluster
// Assumption: optional containers are deployed in nodes that are labelled as optional
// These nodes do not contain mandatory containers
func (nisp NISP) ExecuteForCluster(upperThresholdPower float64) map[string]int32 {
	log.Println("Executing NISP Policy for the entire cluster")
	sortedNodes := kubernetesCluster.GetNodesSortedCPUUsageAscending(constants.OPTIONAL)
	allNodes := kubernetesCluster.GetAllNodeNames()
	return nisp.executePolicy(allNodes, sortedNodes, upperThresholdPower)
}

func (nisp NISP) executePolicy(allNodes []string, sortedNodes []string, upperThresholdPower float64) map[string]int32 {

	i := 0
	predictedPower := powerModel.GetPowerModel().GetPowerConsumptionNodes(allNodes)
	log.Println("Predicted Power", predictedPower)
	if predictedPower < upperThresholdPower {
		log.Println("Predicted power less than upper threshold. Deactivating pods is not possible.")
		return nodeDeployments
	}

	for predictedPower > upperThresholdPower {
		log.Println("===============================================================")
		i++
		log.Println("i: ", i)
		predictedClusterNodes := util.SliceDifference(allNodes, sortedNodes[0:i]) // get the nodes remaining in the cluster after deactivating i nodes

		// get power consumption of the nodes
		predictedPower = powerModel.GetPowerModel().GetPowerConsumptionNodes(predictedClusterNodes)
		log.Println("Predicted Power", predictedPower)
		log.Println("Upper Threshold", upperThresholdPower)
	}

	if (upperThresholdPower-predictedPower)/upperThresholdPower < 0.05 {
		log.Printf("Exact node count used. Deactivating all pods in %v nodes", i)
		nisp.deactivateNodes(sortedNodes[0:i]) // deactivate all pods of 0 to i hosts
		return nodeDeployments
	}

	var policy IPolicyPods = LUCF{} // Can set policy (LUCF, LRU, RCSP)

	if i == 1 {
		log.Printf("i = 1. Executing LUCF for 1st node")
		return policy.ExecuteForNode(sortedNodes[0], upperThresholdPower) // deactivate some containers of 0th node according to a pod selection policy
	} else {
		log.Printf("i = %v. Deactivating all pods in %v nodes", i, i-1)
		nisp.deactivateNodes(sortedNodes[0 : i-1])
		log.Printf("Executing LUCF for %vth node", i)                                    //  node_deployments                             // deactivate all containers of 0 to i-1 hosts
		oneNodeDeployments := policy.ExecuteForNode(sortedNodes[i], upperThresholdPower) // deactivate some containers of ith node according to a pod selection policy
		addToNodeDeployments(oneNodeDeployments)

		return nodeDeployments
	}
}

func (nisp NISP) deactivateNodes(nodeList []string) {
	for _, node := range nodeList {
		oneNodeDeployments := kubernetesCluster.DeactivateNode(node, constants.NAMESPACE, constants.OPTIONAL)
		addToNodeDeployments(oneNodeDeployments)
	}
}

// functions appends values from the map oneNodeDeployments to the map nodeDeployments
func addToNodeDeployments(oneNodeDeployments map[string]int32) {
	for key, value1 := range oneNodeDeployments {
		if value2, exists := nodeDeployments[key]; exists {
			nodeDeployments[key] = value2 + value1
		} else {
			nodeDeployments[key] = value1
		}
	}
}
