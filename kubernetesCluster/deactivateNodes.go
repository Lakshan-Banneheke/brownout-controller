package kubernetesCluster

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
)

// DeactivateNode function returns the deactivated deployment map
func DeactivateNode(nodeName string, namespace string, categoryLabel string) map[string]int32 {
	// cordon the node so that new pods cannot be scheduled on it
	cordonNode(nodeName)

	pods := GetPodsInNodeCategory(nodeName, namespace, categoryLabel)
	deployments := DeactivatePods(pods, namespace)

	log.Printf("All pods in node %s has been deactivated. The node is now idle\n", nodeName)

	return deployments
}

func cordonNode(nodeName string) {
	clientset := getKubernetesClientSet()
	node, err := clientset.CoreV1().Nodes().Get(context.Background(), nodeName, metav1.GetOptions{})
	if err != nil {
		log.Println(err.Error())
	}

	node.Spec.Unschedulable = true
	node, err = clientset.CoreV1().Nodes().Update(context.Background(), node, metav1.UpdateOptions{})
	if err != nil {
		log.Println(err.Error())
	}
	if node.Spec.Unschedulable {
		log.Printf("Node %s has been cordoned\n", nodeName)
	} else {
		log.Printf("Error: Node %s is not cordoned\n", nodeName)
	}
}
