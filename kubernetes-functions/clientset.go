package kubernetes_functions

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"os"
	"path/filepath"
)

func GetKubernetesClientSet() *kubernetes.Clientset {
	var config *rest.Config
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

	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	return clientset
}
