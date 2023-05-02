package kubernetesCluster

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
)

func GetPodNamesAll(namespace string) []string {
	clientset := getKubernetesClientSet()
	// get the list of pods that match the categoryLabel selector (optional or mandatory)
	podList, err := clientset.CoreV1().Pods(namespace).List(context.Background(),
		metav1.ListOptions{})

	if err != nil {
		log.Println(err.Error())
	}

	// create a list of pod names
	var podNames []string
	for _, pod := range podList.Items {
		podNames = append(podNames, pod.Name)
	}

	return podNames
}

func GetTerminatingPodNamesAll(namespace string) []string {
	clientset := getKubernetesClientSet()
	// get the list of pods that match the categoryLabel selector (optional or mandatory)
	podList, err := clientset.CoreV1().Pods(namespace).List(context.Background(),
		metav1.ListOptions{FieldSelector: "status.phase=Terminating"})

	if err != nil {
		log.Println(err.Error())
	}

	// create a list of pod names
	var podNames []string
	for _, pod := range podList.Items {
		podNames = append(podNames, pod.Name)
	}

	return podNames
}

func GetPodNamesCategory(namespace string, categoryLabel string) []string {
	clientset := getKubernetesClientSet()
	// get the list of pods that match the categoryLabel selector (optional or mandatory)
	podList, err := clientset.CoreV1().Pods(namespace).List(context.Background(),
		metav1.ListOptions{LabelSelector: "category=" + categoryLabel})

	if err != nil {
		log.Println(err.Error())
	}

	// create a list of pod names
	var podNames []string
	for _, pod := range podList.Items {
		podNames = append(podNames, pod.Name)
	}

	return podNames
}

func GetPodsInNodeCategory(nodeName string, namespace string, categoryLabel string) []string {
	clientset := getKubernetesClientSet()
	// get the list of pods that match the categoryLabel selector (optional or mandatory)
	podList, err := clientset.CoreV1().Pods(namespace).List(context.Background(),
		metav1.ListOptions{LabelSelector: "category=" + categoryLabel, FieldSelector: "spec.nodeName=" + nodeName})

	if err != nil {
		log.Println(err.Error())
	}

	// create a list of pod names
	var podNames []string
	for _, pod := range podList.Items {
		podNames = append(podNames, pod.Name)
	}

	return podNames
}

func GetPodsInNode(nodeName string, namespace string) []string {
	clientset := getKubernetesClientSet()
	// get the list of pods that match the categoryLabel selector (optional or mandatory)
	podList, err := clientset.CoreV1().Pods(namespace).List(context.Background(),
		metav1.ListOptions{FieldSelector: "spec.nodeName=" + nodeName})

	if err != nil {
		log.Println(err.Error())
	}

	// create a list of pod names
	var podNames []string
	for _, pod := range podList.Items {
		podNames = append(podNames, pod.Name)
	}

	return podNames
}

// GetPodsInNodes : function to retrieve the list of pods in a given set of nodes
func GetPodsInNodes(nodeNames []string, namespace string) []string {
	clientset := getKubernetesClientSet()

	// create a slice of pod names
	var podNames []string

	for _, node := range nodeNames {
		// get the list of pods that resides in the node
		podList, err := clientset.CoreV1().Pods(namespace).List(context.Background(),
			metav1.ListOptions{FieldSelector: "spec.nodeName=" + node})

		if err != nil {
			fmt.Println(err.Error())
		}

		for _, pod := range podList.Items {
			podNames = append(podNames, pod.Name)
		}
	}

	return podNames
}
