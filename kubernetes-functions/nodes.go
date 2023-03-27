package kubernetes_functions

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"
	metrics "k8s.io/metrics/pkg/client/clientset/versioned"
	"strconv"
)

func GetNodeNames(clientset *kubernetes.Clientset, label string) []string {

	labelSelector := labels.SelectorFromSet(labels.Set{"category": label})

	// get the list of nodes that match the label selector
	nodeList, err := clientset.CoreV1().Nodes().List(context.Background(),
		metav1.ListOptions{LabelSelector: labelSelector.String()})

	if err != nil {
		panic(err.Error())
	}

	// create a list of node names
	var nodeNames []string
	for _, node := range nodeList.Items {
		nodeNames = append(nodeNames, node.Name)
	}

	return nodeNames
}

func GetNodesSortedCPUUsage(metricsClient *metrics.Clientset, label string) map[string]int {

	labelSelector := labels.SelectorFromSet(labels.Set{"category": label})

	// get the CPU usage for the node that matches the label selector
	nodeMetrics, err := metricsClient.MetricsV1beta1().NodeMetricses().List(context.Background(),
		metav1.ListOptions{LabelSelector: labelSelector.String()})

	if err != nil {
		panic(err.Error())
	}

	//Make a map of node Name and cpu usage
	nodes := map[string]int{}

	//TODO For testing, Delete later
	nodes["a"] = 12345
	nodes["b"] = 4565465
	nodes["c"] = 2342342344234432
	for _, nodeMetric := range nodeMetrics.Items {
		cpuUsage := nodeMetric.Usage.Cpu().String()
		// Removing the unit "n" from CPU Usage for ease of sorting
		cpuUsageTrimmed := cpuUsage[:len(cpuUsage)-1]
		cpuUsageInt, _ := strconv.Atoi(cpuUsageTrimmed)
		nodes[nodeMetric.ObjectMeta.Name] = cpuUsageInt
	}

	return nodes
}
