package main

import "fmt"

func main() {
	clientset := getKubernetesClientSet()

	nodeNames := getNodesWithLabel(clientset, "optional")
	fmt.Println(nodeNames)
}
