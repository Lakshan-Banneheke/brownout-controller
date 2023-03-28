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

	podNames := kubernetes_functions.GetPodNames(kubernetesClientset, "default", "optional")
	podsCPU := kubernetes_functions.GetPodsSortedCPUUsage(metricsClientSet, "default", "optional")
	//fmt.Println(nodeNames)
	//fmt.Println(nodesCPU)
	fmt.Println(podNames)
	fmt.Println(podsCPU)

}
