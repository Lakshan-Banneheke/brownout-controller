package main

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
	clientset := getKubernetesClientSet()

	nsList, _ := clientset.CoreV1().
		Namespaces().
		List(context.Background(), metav1.ListOptions{})

	for _, n := range nsList.Items {
		fmt.Println(n.Name)
	}
}
