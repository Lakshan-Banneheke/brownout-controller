package kubernetes_functions

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"
	metrics "k8s.io/metrics/pkg/client/clientset/versioned"
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

func GetNodeMetrics(metricsClient *metrics.Clientset, label string) {

	labelSelector := labels.SelectorFromSet(labels.Set{"category": label})

	// get the CPU usage for the node that matches the label selector
	nodeMetrics, err := metricsClient.MetricsV1beta1().NodeMetricses().List(context.Background(),
		metav1.ListOptions{LabelSelector: labelSelector.String()})

	if err != nil {
		panic(err.Error())
	}

	//// create a list of node names
	//var nodeCPUUsages []string
	//for _, nodeMetric := range nodeMetrics.Items {
	//	nodeCPUUsages = append(nodeCPUUsages, node.Name)
	//}
	//
	//return nodeNames

	// print the CPU usage for the node
	for _, nodeMetric := range nodeMetrics.Items {
		fmt.Printf("Node: %s, CPU Usage: %v\n", nodeMetric.ObjectMeta.Name, nodeMetric.Usage.Cpu())
	}
}
