package main

import (
	"brownout-controller/kubernetesCluster"
	"fmt"
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

	deactivatedPodList := []string{"nginx-cd55c47f5-abcsd"}
	deactivatedPodList = kubernetesCluster.DeactivatePods(kubernetesClientset, []string{"nginx-cd55c47f5-k2tt2", "nginx-cd55c47f5-5jsnf", "nginx-cd55c47f5-zprwk", "nginx-cd55c47f5-j56td"}, deactivatedPodList, "default")
	fmt.Println(deactivatedPodList)
}
