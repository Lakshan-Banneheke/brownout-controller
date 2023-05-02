package policies

import (
	"brownout-controller/constants"
	"brownout-controller/kubernetesCluster"
	"brownout-controller/powerModel"
	"brownout-controller/util"
	"log"
)

// NISP implements the IPolicyNodes interface. Node Idling Selection Policy with no migration
type NISP struct {
	deactivatedNodeDeployments map[string]int32
}

// ExecuteForCluster
// Assumption: optional containers are deployed in nodes that are labelled as optional
// These nodes do not contain mandatory containers
func (nisp NISP) ExecuteForCluster(upperThresholdPower float64) map[string]int32 {
	log.Println("Executing NISP Policy for the entire cluster")
	nisp.deactivatedNodeDeployments = make(map[string]int32)
	sortedNodes := kubernetesCluster.GetNodesSortedCPUUsageAscending(constants.OPTIONAL)
	allNodes := kubernetesCluster.GetAllNodeNames()
	return nisp.executePolicy(allNodes, sortedNodes, upperThresholdPower)
}

// function returns a map of deployments scaled down and nodes deactivated
func (nisp NISP) executePolicy(allNodes []string, sortedNodes []string, upperThresholdPower float64) map[string]int32 {

	i := 0
	predictedPower := powerModel.GetPowerModel().GetPowerConsumptionNodes(allNodes)
	log.Println("Current Power", predictedPower)
	if predictedPower < upperThresholdPower {
		log.Println("Current power less than upper threshold. Deactivating pods is not possible.")
		return nisp.deactivatedNodeDeployments
	}

	for predictedPower > upperThresholdPower && i < len(sortedNodes) {
		i++
		predictedClusterNodes := util.SliceDifference(allNodes, sortedNodes[0:i]) // get the nodes remaining in the cluster after deactivating nodes 0 to i (i not inclusive)
		predictedPower = powerModel.GetPowerModel().GetPowerConsumptionNodes(predictedClusterNodes)

		log.Println("===============================================================")
		log.Println("i: ", i)
		log.Println("Predicted Power", predictedPower)
		log.Println("Upper Threshold", upperThresholdPower)
	}

	if (upperThresholdPower-predictedPower)/upperThresholdPower < 0.05 {
		log.Printf("Exact node count used. Deactivating all pods in %v nodes", i)
		nisp.deactivateNodes(sortedNodes[0:i]) // deactivate all pods of nodes 0 to i (i not inclusive)
		return nisp.deactivatedNodeDeployments
	}

	if i == 1 {
		log.Printf("Selected value for i = %v.", i)
		nisp.executePolicyForNode(sortedNodes[i-1], upperThresholdPower)
		return nisp.deactivatedNodeDeployments
	} else {
		log.Printf("Selected value for i = %v. Deactivating all pods in %v nodes", i, i-1)
		nisp.deactivateNodes(sortedNodes[0 : i-1]) // deactivate all containers of nodes 0 to i-1 hosts (i-1 not inclusive)
		nisp.executePolicyForNode(sortedNodes[i-1], upperThresholdPower)
		return nisp.deactivatedNodeDeployments
	}
}

func (nisp NISP) executePolicyForNode(nodeName string, upperThresholdPower float64) {
	policy := GetSelectedPodPolicy(constants.NISP_PER_NODE_POLICY)
	log.Printf("Executing LUCF in node %s", nodeName)
	oneNodeDeactivatedDeployments := policy.ExecuteForNode(nodeName, upperThresholdPower)                                 // deactivate some containers of 0th node according to a pod selection policy
	nisp.deactivatedNodeDeployments = util.AddDeployments(oneNodeDeactivatedDeployments, nisp.deactivatedNodeDeployments) // add to global variable deactivatedNodeDeployments
}

func (nisp NISP) deactivateNodes(nodeList []string) {
	for _, node := range nodeList {
		log.Printf("Deactivating node %s", node)
		oneNodeDeactivatedDeployments := kubernetesCluster.DeactivateNode(node, constants.NAMESPACE)
		nisp.deactivatedNodeDeployments = util.AddDeployments(oneNodeDeactivatedDeployments, nisp.deactivatedNodeDeployments)
	}
}
