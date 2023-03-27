package main

func main() {
	clientset := getKubernetesClientSet()

	getNodesWithLabel(clientset, "optional")
}
