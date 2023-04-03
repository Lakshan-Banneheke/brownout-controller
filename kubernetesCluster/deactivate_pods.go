package kubernetesCluster

import (
	"context"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func DeactivatePod(clientset *kubernetes.Clientset, podName string, namespace string) {
	annotatePodForDeletion(clientset, podName, namespace)

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

	return *pod
}
