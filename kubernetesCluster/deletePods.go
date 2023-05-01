package kubernetesCluster

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func DeletePodsInNode(nodeName string, namespace string, category string) {
	clientset := getKubernetesClientSet()
	podNames := GetPodsInNodeCategory(nodeName, namespace, category)

	for _, pod := range podNames {
		err := clientset.CoreV1().Pods(namespace).Delete(context.TODO(), pod, metav1.DeleteOptions{})
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}
