package main

import (
	"brownout-controller/powerModel"
	"log"
)

func main() {

	//Nodes need to be given the label category=optional via kubectl label nodes <your-node-name> category=optional
	//nodeNames := kubernetesCluster.GetNodeNames("optional")
	//nodesCPU := kubernetesCluster.GetNodesSortedCPUUsage("optional")
	//
	//fmt.Println("==============Nodes ===================")
	//fmt.Println(nodeNames)
	//fmt.Println(nodesCPU)
	//
	//podNames := kubernetesCluster.GetPodNames("default", "optional")
	//podsCPUSorted := kubernetesCluster.GetPodsSortedCPUUsageAll("default", "optional")
	//
	//fmt.Println("==============All Pods ===================")
	//fmt.Println(podNames)
	//fmt.Println(podsCPUSorted)
	//
	//fmt.Println("==============Instance 4===================")
	//podsCPUSortedInstance4 := kubernetesCluster.GetPodsSortedCPUUsageInNode("instance-4", "default", "optional")
	//fmt.Println(podsCPUSortedInstance4)
	//fmt.Println("==============Instance 5===================")
	//podsCPUSortedInstance5 := kubernetesCluster.GetPodsSortedCPUUsageInNode("instance-5", "default", "optional")
	//fmt.Println(podsCPUSortedInstance5)
	//
	//deactivatedDeploymentMap := kubernetesCluster.DeactivatePods([]string{
	//	"nginx-cd55c47f5-jlhq8", "nginx-cd55c47f5-kkqtc",
	//	"traefik-7c57d8789b-dn29b", "traefik-7c57d8789b-pp824"}, "default")
	//
	//fmt.Println(deactivatedDeploymentMap)
	//
	//deployments := map[string]int32{"nginx": 10, "traefik": 5}
	//kubernetesCluster.ActivatePods(deployments, "default")

	//policies.LUCF{}.ExecuteForCluster()

	//lucf := policies.LUCF{}
	//lucf.ExecuteForCluster()

	//var t policies.IPolicy
	//t = policies.LUCF{}
	//t.ExecuteForCluster()

	// get the power model
	pm := powerModel.GetPowerModel()

	// get power consumption when a set of pods given
	log.Println(pm.GetPowerConsumptionPods([]string{"agri-app-master-75656cf88b-fcd29", "agri-app-master-75656cf88b-xtkl4", "agri-app-master-75656cf88b-hxplj"}))
	// get power consumption when a set of nodes given
	log.Println(pm.GetPowerConsumptionNodes([]string{"node-master", "node-worker-1"}))
}
