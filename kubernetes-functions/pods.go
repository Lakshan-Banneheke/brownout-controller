package kubernetes_functions

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/metrics/pkg/apis/metrics/v1beta1"
	metrics "k8s.io/metrics/pkg/client/clientset/versioned"
	"sort"
	"strconv"
)

func GetPodNames(clientset *kubernetes.Clientset, namespace string, categoryLabel string) []string {

	// get the list of pods that match the categoryLabel selector (optional or mandatory)
	podList, err := clientset.CoreV1().Pods(namespace).List(context.Background(),
		metav1.ListOptions{LabelSelector: "category=" + categoryLabel})

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

func GetPodsInNode(nodeName string, clientset *kubernetes.Clientset, namespace string, categoryLabel string) []string {

	// get the list of pods that match the categoryLabel selector (optional or mandatory)
	podList, err := clientset.CoreV1().Pods(namespace).List(context.Background(),
		metav1.ListOptions{LabelSelector: "category=" + categoryLabel, FieldSelector: "spec.nodeName=" + nodeName})

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

	// get the CPU usage for the pod that matches the label selector
	podMetrics, err := metricsClient.MetricsV1beta1().PodMetricses(namespace).List(context.Background(),
		metav1.ListOptions{LabelSelector: "category=" + categoryLabel})

	podsCPUUsage, podNames := extractMetrics(podMetrics.Items, err)

	return sortPodsUsage(podsCPUUsage, podNames)

}

func GetPodsSortedCPUUsageInNode(nodeName string, clientset *kubernetes.Clientset, metricsClient *metrics.Clientset, namespace string, categoryLabel string) []string {

	// get the pods in the given node of the given category
	pods := GetPodsInNode(nodeName, clientset, namespace, categoryLabel)

	// get pod Metrics for all the pods in that node
	var podMetricsItems []v1beta1.PodMetrics
	for _, podName := range pods {
		podMetrics, err := metricsClient.MetricsV1beta1().PodMetricses(namespace).Get(context.Background(), podName, metav1.GetOptions{})
		if err != nil {
			panic(err.Error())
		}
		podMetricsItems = append(podMetricsItems, *podMetrics)
	}

	podsCPUUsage, podNames := extractMetrics(podMetricsItems, nil)

	return sortPodsUsage(podsCPUUsage, podNames)
}

func extractMetrics(podMetricsItems []v1beta1.PodMetrics, err error) (map[string]int, []string) {
	if err != nil {
		panic(err.Error())
	}

	// make a map of pod Name and cpu usage
	podsCPUUsage := map[string]int{}
	var podNames []string

	for _, podMetric := range podMetricsItems {
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

	return podsCPUUsage, podNames
}

// function returns a list of node names in sorted order of increasing cpu usage
func sortPodsUsage(podsCPUUsage map[string]int, podNames []string) []string {

	sort.SliceStable(podNames, func(i, j int) bool {
		return podsCPUUsage[podNames[i]] < podsCPUUsage[podNames[j]]
	})

	return podNames
}
