package prometheus

import (
	"fmt"
	"github.com/prometheus/common/model"
	"log"
)

// interval parameter can have 1m, 30m, 1d, etc
func getTotalRequestCount(hostname string, interval string) int {
	query := fmt.Sprintf("sum by (host) (increase(nginx_ingress_controller_requests{host=~'%s'}[%s]))", hostname, interval)
	result := doQuery(query) // Result is of type Vector
	if result.(model.Vector).Len() == 0 {
		return 0
	}
	vectorVal := result.(model.Vector)[0] // Cast to mode.Vector and get the first row
	reqCount := int(vectorVal.Value)
	log.Printf("Total Request Count for host %s in the last %s: %v", hostname, interval, reqCount)
	return reqCount
}

// interval parameter can have 1m, 30m, 1d, etc
func getErrorRequestCount(hostname string, interval string) int {
	query := fmt.Sprintf("sum by (host) (increase(nginx_ingress_controller_requests{status=~'[4-5].*', host=~'%s'}[%s]))", hostname, interval)
	result := doQuery(query) // Result is of type Vector
	if result.(model.Vector).Len() == 0 {
		return 0
	}
	vectorVal := result.(model.Vector)[0] // Cast to mode.Vector and get the first row
	reqCount := int(vectorVal.Value)
	log.Printf("Error Request Count for host %s in the last %s: %v", hostname, interval, reqCount)
	return reqCount
}

// getSlowRequestCount returns the count of requests which are slower than the given latency
// parameter latency (in seconds): 0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1, 2.5, 5, 10
// parameter interval: 1m, 30m, 1d, etc
func getSlowRequestCount(hostname string, interval string, latency string) int {
	totalSuccessReq := getTotalSuccessRequestCount(hostname, interval)
	fastReq := getFastRequestCount(hostname, interval, latency)
	slowReq := totalSuccessReq - fastReq
	log.Printf("Slow Request Count for host %s in the last %s for latency %s second: %v", hostname, interval, latency, slowReq)
	return slowReq
}

// interval parameter can have 1m, 30m, 1d, etc
func getTotalSuccessRequestCount(hostname string, interval string) int {
	query := fmt.Sprintf("sum by (host) (increase(nginx_ingress_controller_requests{status!~'[4-5].*', host=~'%s'}[%s]))", hostname, interval)
	result := doQuery(query) // Result is of type Vector
	if result.(model.Vector).Len() == 0 {
		return 0
	}
	vectorVal := result.(model.Vector)[0] // Cast to mode.Vector and get the first row
	reqCount := int(vectorVal.Value)
	log.Printf("Total Successful Request Count for host %s in the last %s: %v", hostname, interval, reqCount)
	return reqCount
}

// getFastRequestCount returns the count of requests which are faster than the given latency
// interval parameter can have 1m, 30m, 1d, etc
func getFastRequestCount(hostname string, interval string, latency string) int {
	query := fmt.Sprintf("sum by (le) (increase(nginx_ingress_controller_request_duration_seconds_bucket{status!~'[4-5].*', host=~'%s', le='%s'}[%s]))", hostname, latency, interval)
	result := doQuery(query) // Result is of type Vector
	if result.(model.Vector).Len() == 0 {
		return 0
	}
	vectorVal := result.(model.Vector)[0] // Cast to mode.Vector and get the first row
	reqCount := int(vectorVal.Value)
	log.Printf("Fast Request Count for host %s in the last %s: %v", hostname, interval, reqCount)
	return reqCount
}
