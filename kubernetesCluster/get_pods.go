package kubernetesCluster

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

// GetPodsInNodes : function to retrieve the list of pods in a given set of nodes
func GetPodsInNodes(nodeNames []string, clientset *kubernetes.Clientset, namespace string) []string {

	// create a slice of pod names
	var podNames []string

	for _, node := range nodeNames {
		// get the list of pods that resides in the node
		podList, err := clientset.CoreV1().Pods(namespace).List(context.Background(),
			metav1.ListOptions{FieldSelector: "spec.nodeName=" + node})

		if err != nil {
			panic(err.Error())
		}

		for _, pod := range podList.Items {
			podNames = append(podNames, pod.Name)
		}
	}

	return podNames
}

func GetPodsSortedCPUUsageAll(metricsClient *metrics.Clientset, namespace string, categoryLabel string) []string {

	// get the CPU usage for the pod that matches the label selector
	podMetrics, err := metricsClient.MetricsV1beta1().PodMetricses(namespace).List(context.Background(),
		metav1.ListOptions{LabelSelector: "category=" + categoryLabel})

	podsCPUUsage, podNames := extractCPUMetrics(podMetrics.Items, err)

	return sortPodsUsage(podsCPUUsage, podNames)

}

func GetPodsSortedCPUUsageInNode(nodeName string, clientset *kubernetes.Clientset, metricsClient *metrics.Clientset, namespace string, categoryLabel string) []string {

	// get the pods in the given node of the given category
	pods := GetPodsInNode(nodeName, clientset, namespace, categoryLabel)

	podMetricsItems := getPodMetrics(pods, metricsClient, namespace)  // get pod Metrics for all the pods in that node
	podsCPUUsage, podNames := extractCPUMetrics(podMetricsItems, nil) // filter CPU Usage metric only

	return sortPodsUsage(podsCPUUsage, podNames)
}

// GetPodsCPUUsageSum : function to retrieve the CPU Usage sum of a set of pods
func GetPodsCPUUsageSum(metricsClient *metrics.Clientset, podNames []string, namespace string) float64 {

	// get pod Metrics for all the pods in that node
	podMetricsItems := getPodMetrics(podNames, metricsClient, namespace) // get pod Metrics for all the pods in that node
	podsCPUUsage, _ := extractCPUMetrics(podMetricsItems, nil)           // filter CPU Usage metric only

	// calculate CPU usage sum
	podsCPUUsageSum := 0.0
	for _, cpuUsage := range podsCPUUsage {
		podsCPUUsageSum += float64(cpuUsage)
	}

	return podsCPUUsageSum
}

// GetPodsMemUsageSum : function to retrieve the Memory Usage sum of a set of pods
func GetPodsMemUsageSum(metricsClient *metrics.Clientset, podNames []string, namespace string) float64 {

	podMetricsItems := getPodMetrics(podNames, metricsClient, namespace) // get pod Metrics for all the pods in that node
	podsMemUsage, _ := extractMemMetrics(podMetricsItems, nil)           // filter Memory Usage metric only

	// calculate Mem usage sum
	podsMemUsageSum := 0.0
	for _, memUsage := range podsMemUsage {
		podsMemUsageSum += memUsage
	}

	return podsMemUsageSum
}

// function to retrieve the all metrics of a given set of pods
func getPodMetrics(pods []string, metricsClient *metrics.Clientset, namespace string) []v1beta1.PodMetrics {
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

// function to retrieve the CPU Usage metrics from a given set of metrics
func extractCPUMetrics(podMetricsItems []v1beta1.PodMetrics, err error) (map[string]int, []string) {

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

// function to retrieve the Memory Usage metrics from a given set of metrics
func extractMemMetrics(podMetricsItems []v1beta1.PodMetrics, err error) (map[string]float64, []string) {

	if err != nil {
		panic(err.Error())
	}

	// make a map of pod Name and memory usage
	podsMemUsage := map[string]float64{}
	var podNames []string

	for _, podMetric := range podMetricsItems {
		podMem := 0.0
		for _, cont := range podMetric.Containers {
			contMem := cont.Usage.Memory().String()
			var contMemTrimmed string
			// Removing the unit "Ki" from Memory Usage and converting to float
			if contMem != "0" {
				contMemTrimmed = contMem[:len(contMem)-2]
			} else {
				contMemTrimmed = contMem
			}
			memUsageFloat, err := strconv.ParseFloat(contMemTrimmed, 64)
			if err != nil {
				panic(err.Error())
			}
			podMem += memUsageFloat
		}

		podsMemUsage[podMetric.ObjectMeta.Name] = podMem
		podNames = append(podNames, podMetric.ObjectMeta.Name)
	}

	return podsMemUsage, podNames
}

// function returns a list of node names in sorted order of increasing cpu usage
func sortPodsUsage(podsCPUUsage map[string]int, podNames []string) []string {

	sort.SliceStable(podNames, func(i, j int) bool {
		return podsCPUUsage[podNames[i]] < podsCPUUsage[podNames[j]]
	})

	return podNames
}
