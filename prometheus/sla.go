package prometheus

import (
	"log"
)

func GetSLAViolationRatio(hostname string, interval string, latency string) float64 {
	reqTotal := getTotalRequestCount(hostname, interval)
	reqError := getErrorRequestCount(hostname, interval)
	reqSlow := getSlowRequestCount(hostname, interval, latency)

	slaViolationRatio := float64(reqSlow+reqError) / float64(reqTotal)
	log.Printf("SLA violation ratio for host %v", slaViolationRatio)
	return slaViolationRatio
}

func GetSLASuccessRatio(hostname string, interval string, latency string) float64 {
	slaSuccessRatio := 1 - GetSLAViolationRatio(hostname, interval, latency)
	log.Printf("SLA Success ratio for host %v", slaSuccessRatio)
	return slaSuccessRatio
}
