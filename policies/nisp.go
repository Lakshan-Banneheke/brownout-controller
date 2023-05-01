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

var deactivatedNodeDeployments = make(map[string]int32)

// ExecuteForCluster
// Assumption: optional containers are deployed in nodes that are labelled as optional
// These nodes do not contain mandatory containers
func (nisp NISP) ExecuteForCluster(upperThresholdPower float64) map[string]int32 {
	log.Println("Executing NISP Policy for the entire cluster")
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
		return deactivatedNodeDeployments
	}

	for predictedPower > upperThresholdPower {
		log.Println("===============================================================")
		i++
		log.Println("i: ", i)
		predictedClusterNodes := util.SliceDifference(allNodes, sortedNodes[0:i]) // get the nodes remaining in the cluster after deactivating nodes 0 to i (i not inclusive)

		predictedPower = powerModel.GetPowerModel().GetPowerConsumptionNodesWithMigration(predictedClusterNodes, 1)
		log.Println("Predicted Power", predictedPower)
		log.Println("Upper Threshold", upperThresholdPower)
	}

	if (upperThresholdPower-predictedPower)/upperThresholdPower < 0.05 {
		log.Printf("Exact node count used. Deactivating all pods in %v nodes", i)
		nisp.deactivateNodes(sortedNodes[0:i]) // deactivate all pods of nodes 0 to i (i not inclusive)
		nisp.migrateNode(sortedNodes[i])       // migrate containers of the ith node to the other available nodes
		return deactivatedNodeDeployments
	}

	var policy IPolicyPods = LUCF{} // Can set policy (LUCF, LRU, RCSP)

	if i == 1 {
		log.Printf("i = 1. Executing LUCF for 1st node")
		kubernetesCluster.CordonNode(sortedNodes[0])                                            // cordoning the node before executing LUCF ensures that this node is not considered in node count in the power prediction inside LUCF since it only counts active nodes
		deactivatedNodeDeployments = policy.ExecuteForNode(sortedNodes[0], upperThresholdPower) // deactivate some containers of 0th node according to a pod selection policy
		nisp.migrateNode(sortedNodes[0])                                                        // migrate containers of the 0th node to the other available nodes
		return deactivatedNodeDeployments
	} else {
		log.Printf("i = %v. Deactivating all pods in %v nodes", i, i-1)
		nisp.deactivateNodes(sortedNodes[0 : i-1]) // deactivate all containers of nodes 0 to i-1 hosts (i-1 not inclusive)
		log.Printf("Executing LUCF for %vth node", i)
		kubernetesCluster.CordonNode(sortedNodes[0])
		oneNodeDeactivatedDeployments := policy.ExecuteForNode(sortedNodes[i-1], upperThresholdPower) // deactivate some containers of i-1th node according to a pod selection policy
		nisp.migrateNode(sortedNodes[i-1])                                                            // migrate containers of the i-1th node to the other available nodes
		deactivatedNodeDeployments = util.AddDeployments(oneNodeDeactivatedDeployments, deactivatedNodeDeployments)
		return deactivatedNodeDeployments
	}
}

func (nisp NISP) deactivateNodes(nodeList []string) {
	for _, node := range nodeList {
		oneNodeDeployments := kubernetesCluster.DeactivateNode(node, constants.NAMESPACE, constants.OPTIONAL)
		deactivatedNodeDeployments = util.AddDeployments(oneNodeDeployments, deactivatedNodeDeployments)
	}
}

func (nisp NISP) migrateNode(nodeName string) {
	log.Printf("Migrating all pods in node %s", nodeName)
	kubernetesCluster.CordonNode(nodeName)
	kubernetesCluster.DeletePodsInNode(nodeName, constants.NAMESPACE, constants.OPTIONAL)
}
