package policies

import (
	"brownout-controller/constants"
	"brownout-controller/kubernetesCluster"
	"log"
)

// LUCF implements the IPolicyPods interface and essentially extends the AbstractPolicy struct
type LUCF struct{ AbstractPolicy }

func (lucf LUCF) ExecuteForCluster(upperThresholdPower float64) map[string]int32 {
	log.Println("Executing LUCF Policy for the entire cluster")
	sortedPods := lucf.sortPodsCluster()
	allClusterPods := sortedPods
	return lucf.executePolicy(allClusterPods, sortedPods, upperThresholdPower)
}

func (lucf LUCF) ExecuteForNode(nodeName string, upperThresholdPower float64) map[string]int32 {
	log.Printf("Executing LUCF Policy for the node %s\n", nodeName)
	sortedPods := lucf.sortPodsNode(nodeName)
	allClusterPods := kubernetesCluster.GetPodNames(constants.NAMESPACE, constants.OPTIONAL)
	return lucf.executePolicy(allClusterPods, sortedPods, upperThresholdPower)
}

func (lucf LUCF) sortPodsCluster() []string {
	return kubernetesCluster.GetPodsSortedCPUUsageAllAscending(constants.NAMESPACE, constants.OPTIONAL)
}

func (lucf LUCF) sortPodsNode(nodeName string) []string {
	return kubernetesCluster.GetPodsSortedCPUUsageInNodeAscending(nodeName, constants.NAMESPACE, constants.OPTIONAL)
}
