package kubernetes_functions

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"
	"k8s.io/metrics/pkg/apis/metrics/v1beta1"
	metrics "k8s.io/metrics/pkg/client/clientset/versioned"
	"sort"
	"strconv"
)

func GetPodNames(clientset *kubernetes.Clientset, namespace string, categoryLabel string) []string {

	labelSelector := labels.SelectorFromSet(labels.Set{"category": categoryLabel})

	// get the list of pods that match the categoryLabel selector (optional or mandatory)
	podList, err := clientset.CoreV1().Pods(namespace).List(context.Background(),
		metav1.ListOptions{LabelSelector: labelSelector.String()})

	if err != nil {
		panic(err.Error())
	}

	// create a list of pod names
	var podNames []string
	for _, pod := range podList.Items {
		podNames = append(podNames, pod.Name)
	}

	return podNames
}

func GetPodsSortedCPUUsageAll(metricsClient *metrics.Clientset, namespace string, categoryLabel string) []string {

	labelSelector := labels.SelectorFromSet(labels.Set{"category": categoryLabel})

	// get the CPU usage for the pod that matches the label selector
	podMetrics, err := metricsClient.MetricsV1beta1().PodMetricses(namespace).List(context.Background(),
		metav1.ListOptions{LabelSelector: labelSelector.String()})

	if err != nil {
		panic(err.Error())
	}

	podsSortedCPU := extractAndSortMetrics(podMetrics)

	return podsSortedCPU
}

func GetPodsSortedCPUUsageNode(nodeName string, metricsClient *metrics.Clientset, namespace string, categoryLabel string) []string {

	labelSelector := labels.SelectorFromSet(labels.Set{"category": categoryLabel})

	// get the CPU usage for the pod that matches the label selector
	podMetrics, err := metricsClient.MetricsV1beta1().PodMetricses(namespace).List(context.Background(),
		metav1.ListOptions{LabelSelector: labelSelector.String(), FieldSelector: fmt.Sprintf("spec.nodeName=%s", nodeName)})

	if err != nil {
		panic(err.Error())
	}

	podsSortedCPU := extractAndSortMetrics(podMetrics)

	return podsSortedCPU
}

func extractAndSortMetrics(podMetrics *v1beta1.PodMetricsList) []string {
	//Make a map of pod Name and cpu usage
	podsCPUUsage := map[string]int{}
	var podNames []string

	for _, podMetric := range podMetrics.Items {
		podCPU := 0
		for _, cont := range podMetric.Containers {
			contCPU := cont.Usage.Cpu().String()
			var contCPUTrimmed string
			// Removing the unit "n" from CPU Usage and converting to int for ease of sorting
			if contCPU != "0" {
				contCPUTrimmed = contCPU[:len(contCPU)-1]
			} else {
				contCPUTrimmed = contCPU
			}
			cpuUsageInt, err := strconv.Atoi(contCPUTrimmed)
			if err != nil {
				panic(err.Error())
			}
			podCPU += cpuUsageInt
		}

		podsCPUUsage[podMetric.ObjectMeta.Name] = podCPU
		podNames = append(podNames, podMetric.ObjectMeta.Name)
	}

	podsSortedCPU := sortPodsUsage(podsCPUUsage, podNames)
	return podsSortedCPU
}

// function returns a list of node names in sorted order of increasing cpu usage
func sortPodsUsage(podsCPUUsage map[string]int, podNames []string) []string {

	sort.SliceStable(podNames, func(i, j int) bool {
		return podsCPUUsage[podNames[i]] < podsCPUUsage[podNames[j]]
	})

	return podNames
}
