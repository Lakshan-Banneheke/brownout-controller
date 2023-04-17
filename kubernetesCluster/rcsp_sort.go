package kubernetesCluster

func GetPodsSortedRandomly(namespace string, categoryLabel string) []string {

	podNames := GetPodNames(namespace, categoryLabel)

	return sortPodsRandomly(podNames)
}

func GetPodsSortedRandomlyInNode(nodeName string, namespace string, categoryLabel string) []string {

	// get the pods in the given node of the given category
	podNames := GetPodsInNode(nodeName, namespace, categoryLabel)

	return sortPodsRandomly(podNames)
}
