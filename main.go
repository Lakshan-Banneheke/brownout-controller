package main

import "brownout-controller/prometheus"

func main() {

	//Nodes need to be given the label category=optional via kubectl label nodes <your-node-name> category=optional
	//nodeNames := kubernetesCluster.GetNodeNames("optional")
	//nodesCPU := kubernetesCluster.GetNodesSortedCPUUsage("optional")
	//
	//log.Println("==============Nodes ===================")
	//log.Println(nodeNames)
	//log.Println(nodesCPU)
	//
	//podNames := kubernetesCluster.GetPodNames("default", "optional")
	//podsCPUSorted := kubernetesCluster.GetPodsSortedCPUUsageAll("default", "optional")
	//
	//log.Println("==============All Pods ===================")
	//log.Println(podNames)
	//log.Println(podsCPUSorted)
	//
	//log.Println("==============Instance 4===================")
	//podsCPUSortedInstance4 := kubernetesCluster.GetPodsSortedCPUUsageInNode("instance-4", "default", "optional")
	//log.Println(podsCPUSortedInstance4)
	//log.Println("==============Instance 5===================")
	//podsCPUSortedInstance5 := kubernetesCluster.GetPodsSortedCPUUsageInNode("instance-5", "default", "optional")
	//log.Println(podsCPUSortedInstance5)
	//
	//deactivatedDeploymentMap := kubernetesCluster.DeactivatePods([]string{
	//	"nginx-cd55c47f5-jlhq8", "nginx-cd55c47f5-kkqtc",
	//	"traefik-7c57d8789b-dn29b", "traefik-7c57d8789b-pp824"}, "default")
	//
	//log.Println(deactivatedDeploymentMap)
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
	//pm := powerModel.GetPowerModel("v4")
	//
	//// get power consumption when a set of pods given
	//log.Println(pm.GetPowerConsumptionPods([]string{"agri-app-master-75656cf88b-kmxvs", "agri-app-master-75656cf88b-rn72n", "agri-app-master-75656cf88b-wtp82"}))
	// get power consumption when a set of nodes given
	//log.Println(pm.GetPowerConsumptionNodes([]string{"test-kubernetes-controller-1"}))

	prometheus.ExampleAPI_query("up")

}
