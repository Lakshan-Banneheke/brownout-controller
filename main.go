package main

import (
	"brownout-controller/kubernetes-functions"
	"fmt"
)

func main() {
	kubernetesClientset, metricsClientSet := kubernetes_functions.GetKubernetesClientSet()

	nodeNames := kubernetes_functions.GetNodeNames(kubernetesClientset, "optional")
	nodesCPU := kubernetes_functions.GetNodesSortedCPUUsage(metricsClientSet, "optional")

	fmt.Println("=====================================================")
	fmt.Println(nodeNames)
	fmt.Println(nodesCPU)

}
