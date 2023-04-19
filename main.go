package main

import "brownout-controller/prometheus"

func main() {

	prometheus.GetSlowRequestCount("podinfo.localdev.me", "1d")

}
