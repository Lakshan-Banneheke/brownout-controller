package kubernetesCluster

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/metrics/pkg/apis/metrics/v1beta1"
	"strconv"
)

func GetMasterNodeUsage() (float64, float64) {
	metricsClient := getMetricsClient()

	// get the CPU usage for the master node
	nodeMetrics, err := metricsClient.MetricsV1beta1().NodeMetricses().List(context.Background(),
		metav1.ListOptions{LabelSelector: "node-role.kubernetes.io/master=true"})

	if err != nil {
		panic(err.Error())
	}

	//Make a map of node Name and cpu usage
	masterCPUUsage := 0.0
	masterMemUsage := 0.0

	for _, nodeMetric := range nodeMetrics.Items {
		cpuUsage := nodeMetric.Usage.Cpu().String()
		// Removing the unit "n" from CPU Usage and converting to int
		cpuUsageInt, err := strconv.Atoi(cpuUsage[:len(cpuUsage)-1])
		if err != nil {
			panic(err.Error())
		}
		masterCPUUsage += float64(cpuUsageInt)

		memUsage := nodeMetric.Usage.Memory().String()
		// Removing the unit "Ki" from Mem Usage and converting to float
		memUsageFloat, err := strconv.ParseFloat(memUsage[:len(memUsage)-2], 64)
		if err != nil {
			panic(err.Error())
		}
		masterMemUsage += memUsageFloat
	}

	return masterCPUUsage, masterMemUsage
}

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

// GetPodsMemUsageSum : function to retrieve the Memory Usage sum of a set of pods
func GetPodsMemUsageSum(podNames []string, namespace string) float64 {
	podMetricsItems := getPodMetrics(podNames, namespace)      // get pod Metrics for all the pods in that node
	podsMemUsage, _ := extractMemMetrics(podMetricsItems, nil) // filter Memory Usage metric only

	// calculate Mem usage sum
	podsMemUsageSum := 0.0
	for _, memUsage := range podsMemUsage {
		podsMemUsageSum += memUsage
	}

	return podsMemUsageSum
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
