package main

import (
	"brownout-controller/constants"
	"brownout-controller/prometheus"
	"fmt"
)

func main() {
	fmt.Println(prometheus.GetSLASuccessRatio("podinfo.localdev.me", "1d", constants.SLA_VIOLATION_LATENCY))
	fmt.Println(prometheus.GetSLASuccessRatio("podinfo.localdev.me", "1d", constants.SLA_VIOLATION_LATENCY))
}
