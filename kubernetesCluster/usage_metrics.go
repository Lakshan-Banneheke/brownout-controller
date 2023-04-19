package kubernetesCluster

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/metrics/pkg/apis/metrics/v1beta1"
)

// GetPodsCPUUsageSum : function to retrieve the CPU Usage sum of a set of pods
func GetPodsCPUUsageSum(podNames []string, namespace string) float64 {
	// get pod Metrics for all the pods in that node
	podMetricsItems := getPodMetrics(podNames, namespace)      // get pod Metrics for all the pods in that node
	podsCPUUsage, _ := extractCPUMetrics(podMetricsItems, nil) // filter CPU Usage metric only

	// calculate CPU usage sum
	podsCPUUsageSum := 0.0
	for _, cpuUsage := range podsCPUUsage {
		podsCPUUsageSum += float64(cpuUsage)
	}

	return podsCPUUsageSum
}

// function to retrieve the all metrics of a given set of pods
func getPodMetrics(pods []string, namespace string) []v1beta1.PodMetrics {
	metricsClient := getMetricsClient()
	var podMetricsItems []v1beta1.PodMetrics
	for _, podName := range pods {
		podMetrics, err := metricsClient.MetricsV1beta1().PodMetricses(namespace).Get(context.Background(), podName, metav1.GetOptions{})
		if err != nil {
			panic(err.Error())
		}
		podMetricsItems = append(podMetricsItems, *podMetrics)
	}
	return podMetricsItems
}
