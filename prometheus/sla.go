package prometheus

func GetSLAViolationRatio(hostname string, interval string, latency string) float64 {
	reqTotal := getTotalRequestCount(hostname, interval)
	reqError := getErrorRequestCount(hostname, interval)
	reqSlow := getSlowRequestCount(hostname, interval, latency)

	slaViolationRatio := float64(reqSlow+reqError) / float64(reqTotal)
	return slaViolationRatio
}

func GetSLASuccessRatio(hostname string, interval string, latency string) float64 {
	return 1 - GetSLAViolationRatio(hostname, interval, latency)
}
