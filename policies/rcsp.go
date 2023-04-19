package policies

import (
	"brownout-controller/constants"
	"brownout-controller/kubernetesCluster"
	"log"
)

// RCSP implements the IPolicyPods interface and essentially extends the AbstractPolicy struct
type RCSP struct{ AbstractPolicy }

func (rcsp RCSP) ExecuteForCluster() {
	log.Println("Executing RCSP Policy for the entire cluster")
	sortedPods := rcsp.sortPodsCluster()
	allClusterPods := sortedPods
	rcsp.executePolicy(allClusterPods, sortedPods)
}

func (rcsp RCSP) ExecuteForNode(nodeName string) {
	log.Printf("Executing RCSP Policy for the node %s\n", nodeName)
	sortedPods := rcsp.sortPodsNode(nodeName)
	allClusterPods := kubernetesCluster.GetPodNames(constants.NAMESPACE, constants.OPTIONAL)
	rcsp.executePolicy(allClusterPods, sortedPods)
}

func (rcsp RCSP) sortPodsCluster() []string {
	return kubernetesCluster.GetPodsSortedRandomly(constants.NAMESPACE, constants.OPTIONAL)
}

func (rcsp RCSP) sortPodsNode(nodeName string) []string {
	return kubernetesCluster.GetPodsSortedRandomlyInNode(nodeName, constants.NAMESPACE, constants.OPTIONAL)
}
