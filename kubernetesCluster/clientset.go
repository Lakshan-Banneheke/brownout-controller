package kubernetesCluster

import (
	"fmt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	metrics "k8s.io/metrics/pkg/client/clientset/versioned"
	"os"
	"path/filepath"
)

var kubernetesClientset *kubernetes.Clientset
var metricsClientSet *metrics.Clientset
var config *rest.Config

//func GetClientSets() (*kubernetes.Clientset, *metrics.Clientset) {
//	config := getConfig()
//	kubernetesClientset := getKubernetesClientSet(config)
//	metricsClientSet := getMetricsClient(config)
//	return kubernetesClientset, metricsClientSet
//}

func getKubernetesClientSet() *kubernetes.Clientset {
	// creates the kubernetes kubernetesClientset
	if kubernetesClientset == nil {
		var err error
		kubernetesClientset, err = kubernetes.NewForConfig(getConfig())
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("New Kubernetes Clientset created")
	}

	return kubernetesClientset
}

func getMetricsClient() *metrics.Clientset {
	// create a new metrics client
	if metricsClientSet == nil {
		var err error
		metricsClientSet, err = metrics.NewForConfig(getConfig())
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("New Metrics Clientset created")
	}

	return metricsClientSet
}

func getConfig() *rest.Config {
	if config == nil {
		var err error

		//check environment variables and gets the value of inCluster
		inCluster, present := os.LookupEnv("inCluster")

		//if inCluster env variable is not set, assuming out of cluster configuration
		if !present {
			inCluster = "false"
		}

		if inCluster == "true" {
			// creates the in-cluster config
			config, err = rest.InClusterConfig()
		} else {
			// creates out of cluster configuration assuming kube config file is in "/home/.kube/config"
			config, err = clientcmd.BuildConfigFromFlags(
				"",
				filepath.Join(homedir.HomeDir(), ".kube/config"),
			)
		}

		if err != nil {
			panic(err.Error())
		}
	}

	return config
}
