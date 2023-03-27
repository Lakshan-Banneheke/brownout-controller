package main

import (
	"brownout-controller/kubernetes-functions"
	"fmt"
)

func main() {
	clientset := kubernetes_functions.GetKubernetesClientSet()

	nodeNames := kubernetes_functions.GetNodesWithLabel(clientset, "optional")
	fmt.Println(nodeNames)
}
