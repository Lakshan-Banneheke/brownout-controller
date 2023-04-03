package main

import (
	"brownout-controller/kubernetesCluster"
)

func main() {
	kubernetesClientset, _ := kubernetesCluster.GetClientSets()

	// Nodes need to be given the label category=optional via kubectl label nodes <your-node-name> category=optional
	//nodeNames := kubernetes_functions.GetNodeNames(kubernetesClientset, "optional")
	//nodesCPU := kubernetes_functions.GetNodesSortedCPUUsage(metricsClientSet, "optional")
	//
	//fmt.Println("==============Nodes ===================")
	//fmt.Println(nodeNames)
	//fmt.Println(nodesCPU)
	//
	//podNames := kubernetes_functions.GetPodNames(kubernetesClientset, "default", "optional")
	//podsCPUSorted := kubernetes_functions.GetPodsSortedCPUUsageAll(metricsClientSet, "default", "optional")
	//
	//fmt.Println("==============All Pods ===================")
	//fmt.Println(podNames)
	//fmt.Println(podsCPUSorted)
	//
	//fmt.Println("==============Instance 4===================")
	//podsCPUSortedInstance4 := kubernetes_functions.GetPodsSortedCPUUsageInNode("instance-4", kubernetesClientset, metricsClientSet, "default", "optional")
	//fmt.Println(podsCPUSortedInstance4)
	//fmt.Println("==============Instance 5===================")
	//podsCPUSortedInstance5 := kubernetes_functions.GetPodsSortedCPUUsageInNode("instance-5", kubernetesClientset, metricsClientSet, "default", "optional")
	//fmt.Println(podsCPUSortedInstance5)

	kubernetesCluster.DeactivatePod(kubernetesClientset, "nginx-cd55c47f5-cf8sh", "default")
}
