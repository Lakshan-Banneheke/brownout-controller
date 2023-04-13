package kubernetesCluster

import (
	"context"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// DeactivatePods function returns the deactivated deployment map
func DeactivatePods(clientset *kubernetes.Clientset, podNames []string, namespace string) map[string]int32 {
	var deployments = make(map[string]int32)

	for _, podName := range podNames {
		pod := annotatePodForDeletion(clientset, podName, namespace)
		deploymentName := getDeployment(clientset, pod, namespace)

		if val, exists := deployments[deploymentName]; exists {
			deployments[deploymentName] = val + 1
		} else {
			deployments[deploymentName] = 1
		}
	}

	for deploymentName, value := range deployments {
		scaleDownDeployment(clientset, deploymentName, value, namespace)
	}

	return deployments
}

//func DeactivatePod(clientset *kubernetes.Clientset, podName string, namespace string) {
//	pod := annotatePodForDeletion(clientset, podName, namespace)
//	deploymentName := getDeployment(clientset, pod, namespace)
//	scaleDownDeployment(clientset, deploymentName, 1, namespace)
//}

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

func scaleDownDeployment(clientset *kubernetes.Clientset, deploymentName string, count int32, namespace string) {
	scale, err := clientset.AppsV1().Deployments(namespace).GetScale(context.Background(), deploymentName, metav1.GetOptions{})
	if err != nil {
		panic(err.Error())
	}

	scale.Spec.Replicas -= count

	updatedScale, err := clientset.AppsV1().Deployments(namespace).UpdateScale(context.Background(), deploymentName, scale, metav1.UpdateOptions{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Deployment %s in namespace %s scaled down to %d replicas\n", deploymentName, namespace, updatedScale.Spec.Replicas)
}
