package policies

import (
	"brownout-controller/constants"
	"brownout-controller/kubernetesCluster"
	"log"
)

// RCSP implements the IPolicyPods interface and essentially extends the AbstractPolicy struct
type RCSP struct{ AbstractPolicy }

func (rcsp RCSP) ExecuteForCluster(upperThresholdPower float64) map[string]int32 {
	log.Println("Executing RCSP Policy for the entire cluster")
	sortedPods := rcsp.sortPodsCluster()
	allClusterPods := kubernetesCluster.GetPodNamesAll(constants.NAMESPACE)
	return rcsp.executePolicy(allClusterPods, sortedPods, upperThresholdPower)
}

func (rcsp RCSP) ExecuteForNode(nodeName string, upperThresholdPower float64) map[string]int32 {
	log.Printf("Executing RCSP Policy for the node %s\n", nodeName)
	sortedPods := rcsp.sortPodsNode(nodeName)
	allClusterPods := kubernetesCluster.GetPodNamesAll(constants.NAMESPACE)
	return rcsp.executePolicy(allClusterPods, sortedPods, upperThresholdPower)
}

func (rcsp RCSP) sortPodsCluster() []string {
	return kubernetesCluster.GetPodsSortedRandomly(constants.NAMESPACE, constants.OPTIONAL)
}

func (rcsp RCSP) sortPodsNode(nodeName string) []string {
	return kubernetesCluster.GetPodsSortedRandomlyInNode(nodeName, constants.NAMESPACE, constants.OPTIONAL)
}
