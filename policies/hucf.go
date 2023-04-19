package policies

import (
	"brownout-controller/constants"
	"brownout-controller/kubernetesCluster"
	"log"
)

// HUCF implements the IPolicyPods interface and essentially extends the AbstractPolicy struct
type HUCF struct{ AbstractPolicy }

func (hucf HUCF) ExecuteForCluster() {
	log.Println("Executing HUCF Policy for the entire cluster")
	sortedPods := hucf.sortPodsCluster()
	allClusterPods := sortedPods
	hucf.executePolicy(allClusterPods, sortedPods)
}

func (hucf HUCF) ExecuteForNode(nodeName string) {
	log.Printf("Executing HUCF Policy for the node %s\n", nodeName)
	sortedPods := hucf.sortPodsNode(nodeName)
	allClusterPods := kubernetesCluster.GetPodNames(constants.NAMESPACE, constants.OPTIONAL)
	hucf.executePolicy(allClusterPods, sortedPods)
}

func (hucf HUCF) sortPodsCluster() []string {
	return kubernetesCluster.GetPodsSortedCPUUsageAllDescending(constants.NAMESPACE, constants.OPTIONAL)
}

func (hucf HUCF) sortPodsNode(nodeName string) []string {
	return kubernetesCluster.GetPodsSortedCPUUsageInNodeDescending(nodeName, constants.NAMESPACE, constants.OPTIONAL)
}
