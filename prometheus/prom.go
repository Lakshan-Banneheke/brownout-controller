package prometheus

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

func ExampleAPI_query(q string) {
	// PROMETHEUS_IP should be set in environment variables
	prometheusIP, present := os.LookupEnv("PROMETHEUS_IP")
	if !present {
		panic("PROMETHEUS_IP should be set in environment variables")
	}

	prometheusURL := "http://" + prometheusIP + ":9090"
	client, err := api.NewClient(api.Config{
		Address: prometheusURL,
	})
	if err != nil {
		log.Printf("Error creating client: %v\n", err)
		os.Exit(1)
	}

	log.Println("Created Prometheus Client")

	v1api := v1.NewAPI(client)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, warnings, err := v1api.Query(ctx, q, time.Now(), v1.WithTimeout(5*time.Second))
	if err != nil {
		log.Printf("Error querying Prometheus: %v\n", err)
		os.Exit(1)
	}
	if len(warnings) > 0 {
		log.Printf("Warnings: %v\n", warnings)
	}
	log.Printf("Result:\n%v\n", result)
}
