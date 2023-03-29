package main

import (
	"brownout-controller/kubernetes-functions"
	"fmt"
)

func main() {
	kubernetesClientset, metricsClientSet := kubernetes_functions.GetClientSets()

	// Nodes need to be given the label category=optional
	nodeNames := kubernetes_functions.GetNodeNames(kubernetesClientset, "optional")
	nodesCPU := kubernetes_functions.GetNodesSortedCPUUsage(metricsClientSet, "optional")

	fmt.Println(nodeNames)
	fmt.Println(nodesCPU)

	//namespace is set as wso2 for testing
	podNames := kubernetes_functions.GetPodNames(kubernetesClientset, "default", "optional")
	podsCPUSorted := kubernetes_functions.GetPodsSortedCPUUsageAll(metricsClientSet, "default", "optional")

	fmt.Println("==============All Pods ===================")
	fmt.Println(podNames)
	fmt.Println(podsCPUSorted)

	//podsCPUSortedInstance4 := kubernetes_functions.GetPodsSortedCPUUsageNode("instance-4", metricsClientSet, "wso2", "optional")
	//podsCPUSortedInstance5 := kubernetes_functions.GetPodsSortedCPUUsageNode("instance-5", metricsClientSet, "wso2", "optional")
	//

	//fmt.Println("==============Instance 4===================")
	//fmt.Println(podsCPUSortedInstance4)
	//fmt.Println("==============Instance 5===================")
	//fmt.Println(podsCPUSortedInstance5)

	//fmt.Println("==============")
	//fmt.Println(kubernetes_functions.GetPodsInNode(kubernetesClientset, "wso2", "optional"))
}
