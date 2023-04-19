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
	client, err := api.NewClient(api.Config{
		Address: "http://10.43.218.49:9090",
	})
	if err != nil {
		log.Printf("Error creating client: %v\n", err)
		os.Exit(1)
	}

	log.Printf("Creatied client: %v\n", err)

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
