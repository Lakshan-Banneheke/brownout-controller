package kubernetesCluster

import (
	"context"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func DeactivatePod(clientset *kubernetes.Clientset, podName string, namespace string) {
	pod := annotatePodForDeletion(clientset, podName, namespace)
	deploymentName := getDeployment(clientset, pod, namespace)
	scaleDownDeployment(clientset, deploymentName, namespace)
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

func getDeployment(clientset *kubernetes.Clientset, pod corev1.Pod, namespace string) string {
	ownerPod := pod.ObjectMeta.OwnerReferences
	replicaSetName := ownerPod[0].Name
	replicaSet, _ := clientset.AppsV1().ReplicaSets(namespace).Get(context.Background(), replicaSetName, metav1.GetOptions{})
	ownerRS := replicaSet.ObjectMeta.OwnerReferences
	deploymentName := ownerRS[0].Name
	return deploymentName
}

func scaleDownDeployment(clientset *kubernetes.Clientset, deploymentName string, namespace string) {
	scale, err := clientset.AppsV1().Deployments(namespace).GetScale(context.Background(), deploymentName, metav1.GetOptions{})
	if err != nil {
		panic(err.Error())
	}

	scale.Spec.Replicas -= 1

	updatedScale, err := clientset.AppsV1().Deployments(namespace).UpdateScale(context.Background(), "nginx", scale, metav1.UpdateOptions{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Deployment %s in namespace %s scaled down to %d replicas\n", deploymentName, namespace, updatedScale.Spec.Replicas)
}
