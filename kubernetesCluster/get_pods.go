package kubernetesCluster

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetPodNames(namespace string, categoryLabel string) []string {
	clientset := getKubernetesClientSet()
	// get the list of pods that match the categoryLabel selector (optional or mandatory)
	podList, err := clientset.CoreV1().Pods(namespace).List(context.Background(),
		metav1.ListOptions{LabelSelector: "category=" + categoryLabel})

	if err != nil {
		fmt.Println(err.Error())
	}

	// create a list of pod names
	var podNames []string
	for _, pod := range podList.Items {
		podNames = append(podNames, pod.Name)
	}

	return podNames
}

func GetPodsInNode(nodeName string, namespace string, categoryLabel string) []string {
	clientset := getKubernetesClientSet()
	// get the list of pods that match the categoryLabel selector (optional or mandatory)
	podList, err := clientset.CoreV1().Pods(namespace).List(context.Background(),
		metav1.ListOptions{LabelSelector: "category=" + categoryLabel, FieldSelector: "spec.nodeName=" + nodeName})

	if err != nil {
		fmt.Println(err.Error())
	}

	// create a list of pod names
	var podNames []string
	for _, pod := range podList.Items {
		podNames = append(podNames, pod.Name)
	}

	return podNames
}
