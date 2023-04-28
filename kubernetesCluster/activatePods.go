package kubernetesCluster

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
)

func ActivatePods(deploymentMap map[string]int32, namespace string) {
	uncordonAllNodes()
	for deploymentName, value := range deploymentMap {
		scaleUpDeployment(deploymentName, value, namespace)
	}
}
func scaleUpDeployment(deploymentName string, count int32, namespace string) {
	log.Println("Updating all deployments to initial replica count")
	clientset := getKubernetesClientSet()
	scale, err := clientset.AppsV1().Deployments(namespace).GetScale(context.Background(), deploymentName, metav1.GetOptions{})
	if err != nil {
		log.Println(err.Error())
	}

	scale.Spec.Replicas += count

	updatedScale, err := clientset.AppsV1().Deployments(namespace).UpdateScale(context.Background(), deploymentName, scale, metav1.UpdateOptions{})
	if err != nil {
		log.Println(err.Error())
	}

	log.Printf("Deployment %s in namespace %s scaled up to %d replicas\n", deploymentName, namespace, updatedScale.Spec.Replicas)
}

func uncordonAllNodes() {
	log.Println("Uncordoning all cordoned nodes")
	clientset := getKubernetesClientSet()
	nodes, err := clientset.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Println(err.Error())
	}
	for _, node := range nodes.Items {
		if node.Spec.Unschedulable {
			node.Spec.Unschedulable = false
			_, err := clientset.CoreV1().Nodes().Update(context.Background(), &node, metav1.UpdateOptions{})
			if err != nil {
				log.Println(err.Error())
			}
			log.Printf("Node %s has been uncordoned\n", node.Name)
		}
	}

}
