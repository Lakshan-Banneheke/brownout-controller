package main

import (
	"brownout-controller/constants"
	"brownout-controller/policies/experimentation"
)

func main() {

	//fmt.Println(prometheus.GetSLAViolationRatio(constants.HOSTNAME, constants.SLA_INTERVAL, constants.SLA_VIOLATION_LATENCY))
	//fmt.Println(prometheus.GetSLASuccessRatio(constants.HOSTNAME, constants.SLA_INTERVAL, constants.SLA_VIOLATION_LATENCY))
	experimentation.LUCFExperiment(constants.REQUIRED_SR)

}
