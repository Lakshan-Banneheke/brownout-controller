package prometheus

import (
	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"log"
	"os"
)

var prometheusClient api.Client

func getPrometheusClient() api.Client {
	if prometheusClient == nil {
		// PROMETHEUS_IP should be set in environment variables
		prometheusIP, present := os.LookupEnv("PROMETHEUS_IP")
		if !present {
			panic("PROMETHEUS_IP should be set in environment variables")
		}

		prometheusURL := "http://" + prometheusIP + ":9090"
		var err error
		prometheusClient, err = api.NewClient(api.Config{
			Address: prometheusURL,
		})
		if err != nil {
			log.Printf("Error creating client: %v\n", err)
			panic(err.Error())
		}
		log.Println("Created Prometheus Client")
	}

	return prometheusClient
}

func getV1API() v1.API {
	client := getPrometheusClient()

	v1api := v1.NewAPI(client)
	return v1api
}
