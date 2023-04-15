package kubernetesCluster

import "fmt"

// DeactivateNode function returns the deactivated deployment map
func DeactivateNode(nodeName string, namespace string, categoryLabel string) map[string]int32 {
	pods := GetPodsInNode(nodeName, namespace, categoryLabel)
	deployments := DeactivatePods(pods, namespace)
	// TODO Do something to prevent new pods from being created in that node.
	// Cordon off the node? But then how can it be reconnected
	fmt.Printf("All pods in node %s has been deactivated. The node is now idle\n", nodeName)
	return deployments
}
