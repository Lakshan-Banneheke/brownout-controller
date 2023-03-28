package main

import (
	"brownout-controller/kubernetes-functions"
	"fmt"
)

func main() {
	kubernetesClientset, metricsClientSet := kubernetes_functions.GetClientSets()

	//nodeNames := kubernetes_functions.GetNodeNames(kubernetesClientset, "optional")
	nodesCPU := kubernetes_functions.GetNodesSortedCPUUsage(metricsClientSet, "optional")

	podNames := kubernetes_functions.GetPodNames(kubernetesClientset, "", "optional")
	//fmt.Println(nodeNames)
	fmt.Println(nodesCPU)
	fmt.Println(podNames)

}
