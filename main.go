package main

import (
	"brownout-controller/policies/experimentation"
	"os"
	"strconv"
)

func main() {

	//fmt.Println(prometheus.GetSLAViolationRatio(constants.HOSTNAME, constants.SLA_INTERVAL, constants.SLA_VIOLATION_LATENCY))
	//fmt.Println(prometheus.GetSLASuccessRatio(constants.HOSTNAME, constants.SLA_INTERVAL, constants.SLA_VIOLATION_LATENCY))
	requiredSR := os.Getenv("REQUIRED_SR")
	requiredSRFloat, _ := strconv.ParseFloat(requiredSR, 32)
	experimentation.LUCFExperiment(requiredSRFloat)

}
