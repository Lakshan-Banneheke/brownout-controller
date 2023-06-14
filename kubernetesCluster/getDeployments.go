package kubernetesCluster

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
)

type Deployment struct {
	Name     string
	Replicas int32
}

func GetDeploymentsAll(namespace string) []Deployment {
	clientset := getKubernetesClientSet()
	deploymentList, err := clientset.AppsV1().Deployments(namespace).List(context.Background(), metav1.ListOptions{})

	if err != nil {
		log.Println(err.Error())
	}

	var deploymentObjList []Deployment

	for _, deployment := range deploymentList.Items {
		deploymentObj := Deployment{
			Name:     deployment.Name,
			Replicas: *deployment.Spec.Replicas,
		}
		deploymentObjList = append(deploymentObjList, deploymentObj)
	}

	return deploymentObjList
}
