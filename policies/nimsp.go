package policies

import (
	"brownout-controller/constants"
	"brownout-controller/kubernetesCluster"
	"brownout-controller/powerModel"
	"brownout-controller/util"
	"log"
)

// NIMSP implements the IPolicyNodes interface. Node Idling Selection Policy with no migration
type NIMSP struct {
	deactivatedNodeDeployments map[string]int32
}

// ExecuteForCluster
// Assumption: optional containers are deployed in nodes that are labelled as optional
// These nodes do not contain mandatory containers
func (nimsp NIMSP) ExecuteForCluster(upperThresholdPower float64) map[string]int32 {
	log.Println("Executing NIMSP Policy for the entire cluster")
	nimsp.deactivatedNodeDeployments = make(map[string]int32)
	sortedNodes := kubernetesCluster.GetNodesSortedCPUUsageAscending(constants.OPTIONAL)
	allNodes := kubernetesCluster.GetAllNodeNames()
	return nimsp.executePolicy(allNodes, sortedNodes, upperThresholdPower)
}

// function returns a map of deployments scaled down and nodes deactivated
func (nimsp NIMSP) executePolicy(allNodes []string, sortedNodes []string, upperThresholdPower float64) map[string]int32 {

	i := 0
	predictedPower := powerModel.GetPowerModel().GetPowerConsumptionNodes(allNodes)
	log.Println("Current Power", predictedPower)
	if predictedPower < upperThresholdPower {
		log.Println("Current power less than upper threshold. Deactivating pods is not possible.")
		return nimsp.deactivatedNodeDeployments
	}

	for predictedPower > upperThresholdPower {
		i++
		predictedClusterNodes := util.SliceDifference(allNodes, sortedNodes[0:i]) // get the nodes remaining in the cluster after deactivating nodes 0 to i (i not inclusive)
		predictedPower = powerModel.GetPowerModel().GetPowerConsumptionNodesWithMigration(predictedClusterNodes, 1)

		log.Println("===============================================================")
		log.Println("i: ", i)
		log.Println("Predicted Power", predictedPower)
		log.Println("Upper Threshold", upperThresholdPower)
	}

	if (upperThresholdPower-predictedPower)/upperThresholdPower < 0.05 {
		log.Printf("Exact node count used. Deactivating all pods in %v nodes", i)
		nimsp.deactivateNodes(sortedNodes[0:i]) // deactivate all pods of nodes 0 to i (i not inclusive)
		nimsp.migrateNode(sortedNodes[i])       // migrate containers of the ith node to the other available nodes
		return nimsp.deactivatedNodeDeployments
	}

	if i == 1 {
		log.Printf("Selected value for i = %v.", i)
		nimsp.executePolicyForNode(sortedNodes[i-1], upperThresholdPower)
		nimsp.migrateNode(sortedNodes[i-1]) // migrate containers of the 0th node to the other available nodes
		return nimsp.deactivatedNodeDeployments
	} else {
		log.Printf("Selected value for i = %v. Deactivating all pods in %v nodes", i, i-1)
		nimsp.deactivateNodes(sortedNodes[0 : i-1]) // deactivate all containers of nodes 0 to i-1 hosts (i-1 not inclusive)
		nimsp.executePolicyForNode(sortedNodes[i-1], upperThresholdPower)
		nimsp.migrateNode(sortedNodes[i-1])
		return nimsp.deactivatedNodeDeployments
	}
}

func (nimsp NIMSP) executePolicyForNode(nodeName string, upperThresholdPower float64) {
	policy := GetSelectedPodPolicy(constants.NISP_PER_NODE_POLICY)
	log.Printf("Executing LUCF in node %s", nodeName)
	kubernetesCluster.CordonNode(nodeName)                                                                                  // cordoning the node before executing LUCF ensures that this node is not considered in node count in the power prediction inside LUCF since it only counts active nodes
	oneNodeDeactivatedDeployments := policy.ExecuteForNode(nodeName, upperThresholdPower)                                   // deactivate some containers of 0th node according to a pod selection policy
	nimsp.deactivatedNodeDeployments = util.AddDeployments(oneNodeDeactivatedDeployments, nimsp.deactivatedNodeDeployments) // add to global variable deactivatedNodeDeployments
}

func (nimsp NIMSP) deactivateNodes(nodeList []string) {
	for _, node := range nodeList {
		log.Printf("Deactivating node %s", node)
		oneNodeDeactivatedDeployments := kubernetesCluster.DeactivateNode(node, constants.NAMESPACE)
		nimsp.deactivatedNodeDeployments = util.AddDeployments(oneNodeDeactivatedDeployments, nimsp.deactivatedNodeDeployments)
	}
}

func (nimsp NIMSP) migrateNode(nodeName string) {
	log.Printf("Migrating all pods in node %s", nodeName)
	kubernetesCluster.CordonNode(nodeName)
	kubernetesCluster.DeletePodsInNode(nodeName, constants.NAMESPACE)
}
