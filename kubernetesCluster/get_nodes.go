package kubernetesCluster

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sort"
	"strconv"
)

func GetNodeNames(categoryLabel string) []string {
	clientset := getKubernetesClientSet()
	// get the list of nodes that match the label selector (optional or mandatory or mixed)
	nodeList, err := clientset.CoreV1().Nodes().List(context.Background(),
		metav1.ListOptions{LabelSelector: "category=" + categoryLabel})

	if err != nil {
		fmt.Println(err.Error())
	}

	// create a list of node names
	var nodeNames []string
	for _, node := range nodeList.Items {
		nodeNames = append(nodeNames, node.Name)
	}

	return nodeNames
}

func GetWorkerNodeCount(clientset *kubernetes.Clientset) int {

	// retrieve all nodes in the cluster
	nodeList, err := clientset.CoreV1().Nodes().List(context.Background(),
		metav1.ListOptions{LabelSelector: "node-role.kubernetes.io/worker=true"})

	if err != nil {
		panic(err.Error())
	}

	// count the number of worker nodes
	workerNodeCount := len(nodeList.Items)

	return workerNodeCount
}

func GetNodesSortedCPUUsage(categoryLabel string) []string {
	metricsClient := getMetricsClient()
	// get the CPU usage for the node that matches the label selector
	nodeMetrics, err := metricsClient.MetricsV1beta1().NodeMetricses().List(context.Background(),
		metav1.ListOptions{LabelSelector: "category=" + categoryLabel})

	if err != nil {
		fmt.Println(err.Error())
	}

	//Make a map of node Name and cpu usage
	nodesCPUUsage := map[string]int{}
	var nodeNames []string

	for _, nodeMetric := range nodeMetrics.Items {
		cpuUsage := nodeMetric.Usage.Cpu().String()
		// Removing the unit "n" from CPU Usage and converting to int for ease of sorting
		cpuUsageInt, err := strconv.Atoi(cpuUsage[:len(cpuUsage)-1])
		if err != nil {
			fmt.Println(err.Error())
		}
		nodesCPUUsage[nodeMetric.ObjectMeta.Name] = cpuUsageInt
		nodeNames = append(nodeNames, nodeMetric.ObjectMeta.Name)
	}

	nodesSortedCPU := sortNodesUsage(nodesCPUUsage, nodeNames)

	return nodesSortedCPU
}

func GetMasterNodeUsage(metricsClient *metrics.Clientset) (float64, float64) {

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

// function returns a list of node names in sorted order of increasing cpu usage
func sortNodesUsage(nodesCPUUsage map[string]int, nodeNames []string) []string {

	sort.SliceStable(nodeNames, func(i, j int) bool {
		return nodesCPUUsage[nodeNames[i]] < nodesCPUUsage[nodeNames[j]]
	})

	return nodeNames
}
