package prometheus

import (
	"fmt"
	"github.com/prometheus/common/model"
	"log"
)

// GetTotalRequestCount interval parameter can have 1m, 30m, 1d, etc
func GetTotalRequestCount(hostname string, interval string) {
	query := "sum by (host) (increase(nginx_ingress_controller_requests{host=~\"" + hostname + "\"}[" + interval + "]))"
	result := doQuery(query)              // Result is of type Vector
	vectorVal := result.(model.Vector)[0] // Cast to mode.Vector and get the first row
	reqCount := int(vectorVal.Value)
	log.Printf("Total Request Count for %s in the last %s: %v", hostname, interval, reqCount)
}

// GetSlowRequestCount interval parameter can have 1m, 30m, 1d, etc
func GetSlowRequestCount(hostname string, interval string) {
	query := "sum by (le) (increase(nginx_ingress_controller_request_duration_seconds_bucket{host=~\"" + hostname + "\"}[" + interval + "]))"
	result := doQuery(query) // Result is of type Vector
	fmt.Println(result)
	//vectorVal := result.(model.Vector)[0] // Cast to mode.Vector and get the first row
	//reqCount := int(vectorVal.Value)
	//log.Printf("Total Request Count for %s in the last %s: %v", hostname, interval, reqCount)
}
