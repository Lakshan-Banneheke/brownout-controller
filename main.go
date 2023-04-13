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

	//deactivatedDeploymentMap := kubernetesCluster.DeactivatePods(kubernetesClientset,
	//	[]string{
	//		"nginx-cd55c47f5-g8zp5", "nginx-cd55c47f5-mn5hq", "nginx-cd55c47f5-htr2f", "nginx-cd55c47f5-pqnpx",
	//		"traefik-7c57d8789b-7666j", "traefik-7c57d8789b-gp6vg",
	//		"helloworld-deployment-68c547667c-vfj4b", "helloworld-deployment-68c547667c-zsz6k", "helloworld-deployment-68c547667c-gzpn5"}, "default")
	//
	//fmt.Println(deactivatedDeploymentMap)

	deployments := map[string]int32{"nginx": 10, "traefik": 5}
	kubernetesCluster.ActivatePods(kubernetesClientset, deployments, "default")
}
