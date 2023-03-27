package main

import (
	"brownout-controller/kubernetes-functions"
	"fmt"
)

func main() {
	kubernetesClientset, metricsClientSet := kubernetes_functions.GetKubernetesClientSet()

	nodeNames := kubernetes_functions.GetNodeNames(kubernetesClientset, "optional")
	kubernetes_functions.GetNodeMetrics(metricsClientSet, "optional")
	fmt.Println(nodeNames)

}
