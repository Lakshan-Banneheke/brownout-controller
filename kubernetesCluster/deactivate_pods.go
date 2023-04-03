package kubernetesCluster

import (
	"context"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func DeactivatePod(clientset *kubernetes.Clientset, podName string, namespace string) {
	_ = annotatePodForDeletion(clientset, podName, namespace)

}

func annotatePodForDeletion(clientset *kubernetes.Clientset, podName string, namespace string) corev1.Pod {
	pod, err := clientset.CoreV1().Pods(namespace).Get(context.Background(), podName, metav1.GetOptions{})
	if err != nil {
		panic(err.Error())
	}

	// set the annotation on the Pod object
	annotations := map[string]string{"controller.kubernetes.io/pod-deletion-cost": "-999"}
	pod.SetAnnotations(annotations)

	// update the Pod object
	pod, err = clientset.CoreV1().Pods("default").Update(context.Background(), pod, metav1.UpdateOptions{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Pod %s annotated with controller.kubernetes.io/pod-deletion-cost:-999 and scheduled for deletion\n", podName)

	return *pod
}

func scaleDownDeployment(clientset *kubernetes.Clientset, deploymentName string, namespace string) {
	scale, err := clientset.AppsV1().Deployments(namespace).GetScale(context.Background(), deploymentName, metav1.GetOptions{})
	if err != nil {
		panic(err.Error())
	}

	scale.Spec.Replicas -= 1

	updatedScale, err := clientset.AppsV1().Deployments("default").UpdateScale(context.Background(), "nginx", scale, metav1.UpdateOptions{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Deployment %s in namespace %s scaled down to %d replicas\n", deploymentName, namespace, updatedScale.Spec.Replicas)
}
