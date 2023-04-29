package kubernetesCluster

import (
	"context"
	"log"
	"strconv"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetNodeNames(categoryLabel string) []string {
	clientset := getKubernetesClientSet()
	// get the list of nodes that match the label selector (optional or mandatory or mixed)
	nodeList, err := clientset.CoreV1().Nodes().List(context.Background(),
		metav1.ListOptions{LabelSelector: "category=" + categoryLabel})

	if err != nil {
		log.Println(err.Error())
	}

	// create a list of node names
	var nodeNames []string
	for _, node := range nodeList.Items {
		nodeNames = append(nodeNames, node.Name)
	}

	return nodeNames
}

func GetAllNodeNames() []string {
	clientset := getKubernetesClientSet()
	// get the list of all nodes (optional & mandatory & mixed)
	nodeList, err := clientset.CoreV1().Nodes().List(context.Background(),
		metav1.ListOptions{})

	if err != nil {
		log.Println(err.Error())
	}

	// create a list of node names
	var nodeNames []string
	for _, node := range nodeList.Items {
		nodeNames = append(nodeNames, node.Name)
	}

	return nodeNames
}

func GetActiveWorkerNodeCount() int {
	clientset := getKubernetesClientSet()
	// retrieve all nodes in the cluster
	nodeList, err := clientset.CoreV1().Nodes().List(context.Background(),
		metav1.ListOptions{LabelSelector: "!node-role.kubernetes.io/master", FieldSelector: "spec.unschedulable=false"})

	if err != nil {
		panic(err.Error())
	}

	// count the number of worker nodes
	workerNodeCount := len(nodeList.Items)

	return workerNodeCount
}

func GetNodesSortedCPUUsageAscending(categoryLabel string) []string {
	metricsClient := getMetricsClient()
	// get the CPU usage for the node that matches the label selector
	nodeMetrics, err := metricsClient.MetricsV1beta1().NodeMetricses().List(context.Background(),
		metav1.ListOptions{LabelSelector: "category=" + categoryLabel})

	if err != nil {
		log.Println(err.Error())
	}

	//Make a map of node Name and cpu usage
	nodesCPUUsage := map[string]int{}
	var nodeNames []string

	for _, nodeMetric := range nodeMetrics.Items {
		cpuUsage := nodeMetric.Usage.Cpu().String()
		// Removing the unit "n" from CPU Usage and converting to int for ease of sorting
		cpuUsageInt, err := strconv.Atoi(cpuUsage[:len(cpuUsage)-1])
		if err != nil {
			log.Println(err.Error())
		}
		nodesCPUUsage[nodeMetric.ObjectMeta.Name] = cpuUsageInt
		nodeNames = append(nodeNames, nodeMetric.ObjectMeta.Name)
	}

	nodesSortedCPU := sortNodesUsageAscending(nodesCPUUsage, nodeNames)

	return nodesSortedCPU
}
