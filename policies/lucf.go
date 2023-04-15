package policies

import (
	"brownout-controller/constants"
	"brownout-controller/kubernetesCluster"
)

// LUCF implements the IPolicy interface and essentially extends the AbstractPolicy struct
type LUCF struct{ AbstractPolicy }

func (lucf LUCF) ExecuteForCluster() {
	sortedPods := lucf.sortPodsCluster()
	allClusterPods := sortedPods
	lucf.executePolicy(allClusterPods, sortedPods)
}

func (lucf LUCF) ExecuteForNode(nodeName string) {
	sortedPods := lucf.sortPodsNode(nodeName)
	allClusterPods := kubernetesCluster.GetPodNames(constants.NAMESPACE, constants.OPTIONAL)
	lucf.executePolicy(allClusterPods, sortedPods)
}

func (lucf LUCF) sortPodsCluster() []string {
	return kubernetesCluster.GetPodsSortedCPUUsageAll(constants.NAMESPACE, constants.OPTIONAL)
}

func (lucf LUCF) sortPodsNode(nodeName string) []string {
	return kubernetesCluster.GetPodsSortedCPUUsageInNode(nodeName, constants.NAMESPACE, constants.OPTIONAL)
}
