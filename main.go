package main

import (
	"brownout-controller/powerModel"
	"log"
)

func main() {

	// get the power model
	pm := powerModel.GetPowerModel("v4")

	// get power consumption when a set of pods given
	log.Println(pm.GetPowerConsumptionPods([]string{"agri-app-master-75656cf88b-kmxvs", "agri-app-master-75656cf88b-rn72n", "agri-app-master-75656cf88b-wtp82"}, "default"))
	// get power consumption when a set of nodes given
	//log.Println(pm.GetPowerConsumptionNodes([]string{"test-kubernetes-controller-1"}, "default"))

}
