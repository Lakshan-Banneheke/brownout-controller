package prometheus

import (
	"context"
	"github.com/prometheus/common/model"
	"log"
	"time"

	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

func doQuery(query string) model.Value {

	v1api := getV1API()

	result, warnings, err := v1api.Query(context.Background(), query, time.Now(), v1.WithTimeout(30*time.Second))

	if err != nil {
		log.Printf("Error querying Prometheus: %v\n", err)
	}
	if len(warnings) > 0 {
		log.Printf("Warnings: %v\n", warnings)
	}

	return result
}
