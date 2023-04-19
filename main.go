package main

import (
	"brownout-controller/brownout"
	"brownout-controller/constants"
	"brownout-controller/powerModel"
	"brownout-controller/prometheus"
	"fmt"
	"log"
)

func main() {

	// get the power model
	pm := powerModel.GetPowerModel()

	// get power consumption when a set of pods given
	log.Println(pm.GetPowerConsumptionPods([]string{"agri-app-master-75656cf88b-fcd29", "agri-app-master-75656cf88b-xtkl4", "agri-app-master-75656cf88b-hxplj"}))
	// get power consumption when a set of nodes given
	log.Println(pm.GetPowerConsumptionNodes([]string{"node-master", "node-worker-1"}))

	fmt.Println(prometheus.GetSLAViolationRatio("podinfo.localdev.me", "1d", constants.SLA_VIOLATION_LATENCY))
	fmt.Println(prometheus.GetSLASuccessRatio("podinfo.localdev.me", "1d", constants.SLA_VIOLATION_LATENCY))

	brownout.ExecuteBrownout()
}
