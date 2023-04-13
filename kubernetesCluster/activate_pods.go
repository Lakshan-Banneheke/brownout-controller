package kubernetesCluster

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ActivatePods(deploymentMap map[string]int32, namespace string) {
	for deploymentName, value := range deploymentMap {
		scaleUpDeployment(deploymentName, value, namespace)
	}
}
func scaleUpDeployment(deploymentName string, count int32, namespace string) {
	clientset := GetKubernetesClientSet()
	scale, err := clientset.AppsV1().Deployments(namespace).GetScale(context.Background(), deploymentName, metav1.GetOptions{})
	if err != nil {
		panic(err.Error())
	}

	scale.Spec.Replicas += count

	updatedScale, err := clientset.AppsV1().Deployments(namespace).UpdateScale(context.Background(), deploymentName, scale, metav1.UpdateOptions{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Deployment %s in namespace %s scaled up to %d replicas\n", deploymentName, namespace, updatedScale.Spec.Replicas)
}
