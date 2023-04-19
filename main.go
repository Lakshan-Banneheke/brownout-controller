package main

import "brownout-controller/prometheus"

func main() {

	prometheus.GetTotalRequestCount("podinfo.localdev.me", "1d")
	prometheus.GetErrorRequestCount("podinfo.localdev.me", "1d")
}
