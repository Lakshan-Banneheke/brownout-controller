package kubernetesCluster

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/metrics/pkg/apis/metrics/v1beta1"
	"log"
)

func GetPodsSortedCPUUsageAllDescending(namespace string, categoryLabel string) []string {
	metricsClient := getMetricsClient()
	// get the CPU usage for the pod that matches the label selector
	podMetrics, err := metricsClient.MetricsV1beta1().PodMetricses(namespace).List(context.Background(),
		metav1.ListOptions{LabelSelector: "category=" + categoryLabel})

	podsCPUUsage, podNames := extractCPUMetrics(podMetrics.Items, err)

	return sortPodsUsageDescending(podsCPUUsage, podNames)

}

func GetPodsSortedCPUUsageInNodeDescending(nodeName string, namespace string, categoryLabel string) []string {
	metricsClient := getMetricsClient()
	// get the pods in the given node of the given category
	pods := GetPodsInNode(nodeName, namespace, categoryLabel)

	// get pod Metrics for all the pods in that node
	var podMetricsItems []v1beta1.PodMetrics
	for _, podName := range pods {
		podMetrics, err := metricsClient.MetricsV1beta1().PodMetricses(namespace).Get(context.Background(), podName, metav1.GetOptions{})
		if err != nil {
			log.Println(err.Error())
		}
		podMetricsItems = append(podMetricsItems, *podMetrics)
	}

	podsCPUUsage, podNames := extractCPUMetrics(podMetricsItems, nil)

	return sortPodsUsageDescending(podsCPUUsage, podNames)
}
