package main

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
}
