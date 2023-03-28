package main

import (
	"brownout-controller/kubernetes-functions"
	"fmt"
)

func main() {
	kubernetesClientset, metricsClientSet := kubernetes_functions.GetClientSets()

	// Nodes need to be given the label category=optional
	//nodeNames := kubernetes_functions.GetNodeNames(kubernetesClientset, "optional")
	//nodesCPU := kubernetes_functions.GetNodesSortedCPUUsage(metricsClientSet, "optional")

	//namespace is set as wso2 for testing
	podNames := kubernetes_functions.GetPodNames(kubernetesClientset, "wso2", "optional")
	podsCPUSorted := kubernetes_functions.GetPodsSortedCPUUsageAll(metricsClientSet, "wso2", "optional")

	podsCPUSortedInstance4 := kubernetes_functions.GetPodsSortedCPUUsageNode("instance-4", metricsClientSet, "wso2", "optional")
	podsCPUSortedInstance5 := kubernetes_functions.GetPodsSortedCPUUsageNode("instance-5", metricsClientSet, "wso2", "optional")

	//fmt.Println(nodeNames)
	//fmt.Println(nodesCPU)
	fmt.Println("==============All Pods ===================")
	fmt.Println(podNames)
	fmt.Println(podsCPUSorted)
	fmt.Println("==============Instance 4===================")
	fmt.Println(podsCPUSortedInstance4)
	fmt.Println("==============Instance 5===================")
	fmt.Println(podsCPUSortedInstance5)
}
