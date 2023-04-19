package main

import "brownout-controller/prometheus"

func main() {

	//prometheus.GetTotalRequestCount("podinfo.localdev.me", "1d")

	//prometheus.GetSlowRequestCount("podinfo.localdev.me", "1d")
	prometheus.GetFastRequestCount("podinfo.localdev.me", "1d")
}
