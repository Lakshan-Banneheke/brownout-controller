package main

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"
)

func getNodesWithLabel(clientset *kubernetes.Clientset, label string) []string {

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
