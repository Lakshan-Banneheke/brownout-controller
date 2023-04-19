package prometheus

import "log"

// GetTotalRequestCount interval parameter can have 1m, 30m, 1d, etc
func GetTotalRequestCount(hostname string, interval string) {
	query := "sum(increase(nginx_ingress_controller_requests{host=~\"" + hostname + "\"}[" + interval + "]))"
	result := doQuery(query)
	log.Printf("Query Result:\n%v\n", result)
}
