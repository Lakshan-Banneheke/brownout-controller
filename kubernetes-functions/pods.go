package kubernetes_functions

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"
)

func GetPodNames(clientset *kubernetes.Clientset, namespace string, label string) []string {

	labelSelector := labels.SelectorFromSet(labels.Set{"category": label})

	// get the list of pods that match the label selector (optional or mandatory)
	podList, err := clientset.CoreV1().Pods(namespace).List(context.Background(),
		metav1.ListOptions{LabelSelector: labelSelector.String()})

	if err != nil {
		panic(err.Error())
	}

	// create a list of node names
	var podNames []string
	for _, pod := range podList.Items {
		podNames = append(podNames, pod.Name)
	}

	return podNames
}
